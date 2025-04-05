package ffmpeg

import (
	"fmt"
	"slices"
	"time"
)

func (f FFmpeg) Trim(id int, from time.Time, length time.Duration) error {
	args := slices.Concat(
		f.getTrimSeek(from),
		f.getTrimInput(),
		f.getTrimLength(length),
		f.getTrimOutput(id),
	)

	return f.execute(args)
}

func (f FFmpeg) getTrimSeek(from time.Time) []string {
	return []string{
		"-ss",
		from.Format("15:04:05.000"),
	}
}

func (f FFmpeg) getTrimInput() []string {
	return []string{
		"-i",
		fmt.Sprintf("./%s/%s.mkv", f.config.Folder.Input, f.config.Parameter.Episode),
	}
}

func (f FFmpeg) getTrimLength(length time.Duration) []string {
	return []string{
		"-t",
		fmt.Sprintf("%.3f", length.Seconds()),
	}
}

func (f FFmpeg) getTrimOutput(id int) []string {
	return []string{
		fmt.Sprintf("./%s/%d.mkv", f.config.Folder.Temporary, id),
	}
}
