package caption

import (
	"fmt"
	"log"
	"oneliner-generator/config"
	"oneliner-generator/filesystem"
	"oneliner-generator/subtitle"
	"os/exec"
)

type FFmpeg struct {
	config     config.Config
	filesystem filesystem.Filesystem
}

func NewFFmpeg(config config.Config, filesystem filesystem.Filesystem) FFmpeg {
	return FFmpeg{
		config:     config,
		filesystem: filesystem,
	}
}

func (f FFmpeg) Run(subtitles []subtitle.Subtitle) {
	for i, subtitle := range subtitles {
		fmt.Printf("%d/%d - %+v\n", i+1, len(subtitles), subtitle)

		f.trim(subtitle)
		f.addSubtitles(subtitle)

		if subtitle.Id == 10 {
			return
		}
	}
}

func (f FFmpeg) trim(subtitle subtitle.Subtitle) {
	args := []string{
		"-ss", subtitle.Duration.From.Format("15:04:05.000"),
		"-i", fmt.Sprintf("./%s/%s.mkv", f.config.Folder.Input, f.config.Parameter.Episode),
		"-t", fmt.Sprintf("%.3f", subtitle.Duration.Length.Seconds()),
		fmt.Sprintf("./%s/%d.mkv", f.config.Folder.Temporary, subtitle.Id),
	}

	cmd := exec.Command("ffmpeg", args...)
	if out, err := cmd.CombinedOutput(); err != nil {
		fmt.Println(string(out))
		log.Fatal(err)
	}
}

func (f FFmpeg) addSubtitles(subtitle subtitle.Subtitle) {
	args := []string{
		"-i", fmt.Sprintf("./%s/%d.mkv", f.config.Folder.Temporary, subtitle.Id),
		"-vf", fmt.Sprintf(
			"subtitles=./%s/%d.srt:force_style='Fontsize=%d',scale=%d:-1:flags=bicubic,fps=%d",
			f.config.Folder.Temporary,
			subtitle.Id,
			f.config.Gif.Subtitle.Size,
			f.config.Gif.Resolution,
			f.config.Gif.Fps,
		),
		fmt.Sprintf("./%s/%s/%s.gif", f.config.Folder.Output, f.config.Parameter.Episode, subtitle.Filename),
	}

	cmd := exec.Command("ffmpeg", args...)
	if out, err := cmd.CombinedOutput(); err != nil {
		fmt.Println(string(out))
		log.Fatal(err)
	}
}
