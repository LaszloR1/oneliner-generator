package ffmpeg

import (
	"fmt"
	"log"
	"oneliner-generator/caption"
	"os/exec"
)

func Trim(mkv_file string, s caption.Subtitle) {
	args := []string{
		"-ss", s.From,
		"-i", mkv_file,
		"-t", s.Duration,
		fmt.Sprintf("_/tmp/%d.mkv", s.Id),
	}

	cmd := exec.Command("ffmpeg", args...)
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}

func AddSubtitles(s caption.Subtitle) {
	args := []string{
		"-i", fmt.Sprintf("_/tmp/%d.mkv", s.Id),
		"-vf", fmt.Sprintf("subtitles=_/tmp/%d.srt:force_style='Fontsize=24',scale=480:-1:flags=bicubic,fps=24", s.Id),
		fmt.Sprintf("_/out/%d. %s.gif", s.Id, s.Filename),
	}

	cmd := exec.Command("ffmpeg", args...)
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}

}
