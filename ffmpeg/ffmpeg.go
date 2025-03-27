package ffmpeg

import (
	"fmt"
	"oneliner-generator/subtitle"
	"os/exec"
)

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
