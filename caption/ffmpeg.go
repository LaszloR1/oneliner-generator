package caption

import (
	"fmt"
	"log"
	"oneliner-generator/types"
	"oneliner-generator/util"
	"os/exec"
)

type FFmpeg struct {
	config     types.Config
	filesystem util.FileSystem
	subtitles  types.Subtitles
	name       string
}

func NewFFmpeg(config types.Config, subtitles types.Subtitles, filesystem util.FileSystem, name string) FFmpeg {
	return FFmpeg{
		config:     config,
		subtitles:  subtitles,
		filesystem: filesystem,
		name:       name,
	}
}

func (f FFmpeg) Run() {
	for i, s := range f.subtitles {
		fmt.Printf("%d/%d - %+v\n", i+1, len(f.subtitles), s)

		f.createTempSrt(s)
		f.trim(f.name, s)
		f.addSubtitles(s)
	}
}

func (f FFmpeg) createTempSrt(subtitle types.Subtitle) {
	name := fmt.Sprintf("%d.srt", subtitle.Id)
	contents := fmt.Sprintf("1\n00:00:00,000 --> 00:01:00,000\n%s\n%s\n", subtitle.Line1, subtitle.Line2)

	f.filesystem.CreateTempFile(name, contents)
}

func (f FFmpeg) trim(name string, s types.Subtitle) {
	args := []string{
		"-ss", s.From,
		"-i", fmt.Sprintf("./%s/%s.mkv", f.config.InputFolder, name),
		"-t", s.Duration,
		fmt.Sprintf("./%s/%d.mkv", f.config.TempFolder, s.Id),
	}

	cmd := exec.Command("ffmpeg", args...)
	if out, err := cmd.CombinedOutput(); err != nil {
		println(string(out))
		log.Fatal(err)
	}
}

func (f FFmpeg) addSubtitles(s types.Subtitle) {
	args := []string{
		"-i", fmt.Sprintf("./%s/%d.mkv", f.config.TempFolder, s.Id),
		"-vf", fmt.Sprintf(
			"subtitles=./%s/%d.srt:force_style='Fontsize=%d',scale=%d:-1:flags=bicubic,fps=%d",
			f.config.TempFolder,
			s.Id, f.config.SubtitleFontsize,
			f.config.GifResolution,
			f.config.GifFramerate,
		),
		fmt.Sprintf("./%s/%d. %s.gif", f.config.OutputFolder, s.Id, s.Filename),
	}

	cmd := exec.Command("ffmpeg", args...)
	if out, err := cmd.CombinedOutput(); err != nil {
		println(string(out))
		log.Fatal(err)
	}
}
