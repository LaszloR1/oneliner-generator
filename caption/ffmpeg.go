package caption

import (
	"fmt"
	"log"
	"oneliner-generator/config"
	"oneliner-generator/types"
	"oneliner-generator/util"
	"os/exec"
)

type FFmpeg struct {
	config     config.Config
	filesystem util.FileSystem
	subtitles  types.Subtitles
	name       string
}

func NewFFmpeg(config config.Config, subtitles types.Subtitles, filesystem util.FileSystem, name string) FFmpeg {
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

		f.trim(f.name, s)
		f.addSubtitles(s)
	}
}

func (f FFmpeg) trim(name string, s types.Subtitle) {
	args := []string{
		"-ss", s.From,
		"-i", fmt.Sprintf("./%s/%s.mkv", f.config.Folder.Input, name),
		"-t", s.Duration,
		fmt.Sprintf("./%s/%d.mkv", f.config.Folder.Temporary, s.Id),
	}

	cmd := exec.Command("ffmpeg", args...)
	if out, err := cmd.CombinedOutput(); err != nil {
		println(string(out))
		log.Fatal(err)
	}
}

func (f FFmpeg) addSubtitles(s types.Subtitle) {
	args := []string{
		"-i", fmt.Sprintf("./%s/%d.mkv", f.config.Folder.Temporary, s.Id),
		"-vf", fmt.Sprintf(
			"subtitles=./%s/%d.srt:force_style='Fontsize=%d',scale=%d:-1:flags=bicubic,fps=%d",
			f.config.Folder.Temporary,
			s.Id,
			f.config.Gif.Subtitle.Size,
			f.config.Gif.Resolution,
			f.config.Gif.Fps,
		),
		fmt.Sprintf("./%s/%s/%s", f.config.Folder.Output, f.name, s.Filename),
	}

	cmd := exec.Command("ffmpeg", args...)
	if out, err := cmd.CombinedOutput(); err != nil {
		println(string(out))
		log.Fatal(err)
	}
}
