package ffmpeg

import (
	"fmt"
	"oneliner-generator/subtitle"
)

func (f FFmpeg) trim(subtitle subtitle.Subtitle) error {
	args := []string{
		"-ss", f.getTrimSeekStr(subtitle),
		"-i", f.getTrimInputStr(),
		"-t", f.getTrimLengthStr(subtitle),
		f.getTrimOutputStr(subtitle),
	}

	return f.execute(args)
}

func (f FFmpeg) getTrimSeekStr(subtitle subtitle.Subtitle) string {
	return subtitle.Duration.From.Format("15:04:05.000")
}

func (f FFmpeg) getTrimInputStr() string {
	return fmt.Sprintf("./%s/%s.mkv", f.config.Folder.Input, f.config.Parameter.Episode)
}

func (f FFmpeg) getTrimLengthStr(subtitle subtitle.Subtitle) string {
	return fmt.Sprintf("%.3f", subtitle.Duration.Length.Seconds())
}

func (f FFmpeg) getTrimOutputStr(subtitle subtitle.Subtitle) string {
	return fmt.Sprintf("./%s/%d.mkv", f.config.Folder.Temporary, subtitle.Id)
}
