package ffmpeg

import (
	"fmt"
	"slices"
)

func (f FFmpeg) Extract() error {
	args := slices.Concat(
		f.getExtractInput(),
		f.getExtractOutput(),
	)

	return f.execute(args)
}

func (f FFmpeg) getExtractInput() []string {
	return []string{
		"-i",
		fmt.Sprintf("./%s/%s.mkv", f.config.Folder.Input, f.config.Parameter.Episode),
	}
}

func (f FFmpeg) getExtractOutput() []string {
	return []string{
		fmt.Sprintf("./%s/%s.srt", f.config.Folder.Input, f.config.Parameter.Episode),
	}
}
