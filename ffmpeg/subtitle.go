package ffmpeg

import (
	"fmt"
	"slices"
)

func (f FFmpeg) AddSubtitles(id int, filename string) error {
	args := slices.Concat(
		f.getSubtitleInput(id),
		f.getSubtitleFilter(id),
		f.getSubtitleOutput(filename),
	)

	return f.execute(args)
}

func (f FFmpeg) getSubtitleInput(id int) []string {
	return []string{
		"-i",
		fmt.Sprintf("./%s/%d.mkv", f.config.Folder.Temporary, id),
	}
}

func (f FFmpeg) getSubtitleFilter(id int) []string {
	return []string{
		"-vf",
		fmt.Sprintf(
			"subtitles=./%s/%d.srt:force_style='FontName=%s,Fontsize=%d',scale=%d:-1:flags=bicubic,fps=%d",
			f.config.Folder.Temporary,
			id,
			f.config.Gif.Subtitle.Font,
			f.config.Gif.Subtitle.Size,
			f.config.Gif.Resolution,
			f.config.Gif.Fps,
		)}
}

func (f FFmpeg) getSubtitleOutput(filename string) []string {
	return []string{
		fmt.Sprintf("./%s/%s/%s.gif", f.config.Folder.Output, f.config.Parameter.Episode, filename),
	}
}
