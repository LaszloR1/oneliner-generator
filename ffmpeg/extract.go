package ffmpeg

import (
	"fmt"
	"slices"
)

func (f FFmpeg) Extract() error {
	args := slices.Concat(
		f.getExtractInput(),
		f.getMap(),
		f.getExtractOutput(),
	)

	return f.execute(args)
}

func (f FFmpeg) getExtractInput() []string {
	return []string{
		"-i",
		fmt.Sprintf("./%s/%s.%s", f.config.Folder.Input, f.config.Parameter.Episode, f.config.Parameter.Format),
	}
}

func (f FFmpeg) getMap() []string {
	return []string{
		"-map",
		fmt.Sprintf("0:s:%d", f.config.Parameter.Subtitle),
	}
}

func (f FFmpeg) getExtractOutput() []string {
	return []string{
		fmt.Sprintf("./%s/%s.srt", f.config.Folder.Input, f.config.Parameter.Episode),
	}
}
