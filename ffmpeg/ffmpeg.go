package ffmpeg

import (
	"fmt"
	"oneliner-generator/config"
	"oneliner-generator/filesystem"
	"oneliner-generator/subtitle"
	"os/exec"
)

type FFmpeg struct {
	config config.Config
	fs     filesystem.Filesystem
}

func New(config config.Config, fs filesystem.Filesystem) FFmpeg {
	return FFmpeg{
		config: config,
		fs:     fs,
	}
}

func (f FFmpeg) Run(subtitles []subtitle.Subtitle) error {
	for i, subtitle := range subtitles {
		fmt.Printf("%d/%d - %+v\n", i+1, len(subtitles), subtitle)

		if err := f.trim(subtitle); err != nil {
			return err
		}

		if err := f.addSubtitles(subtitle); err != nil {
			return err
		}
	}

	return nil
}

func (f FFmpeg) execute(args []string) error {
	cmd := exec.Command("ffmpeg", args...)
	if out, err := cmd.CombinedOutput(); err != nil {
		fmt.Println(string(out))
		return err
	}

	return nil

}
