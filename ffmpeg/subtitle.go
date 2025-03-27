package ffmpeg

import (
	"fmt"
	"oneliner-generator/subtitle"
)

func (f FFmpeg) addSubtitles(subtitle subtitle.Subtitle) error {
	args := []string{
		"-i", f.getSubtitleInputStr(subtitle),
		"-vf", f.getSubtitleFilterStr(subtitle),
		f.getSubtitleOutputStr(subtitle),
	}

	return f.execute(args)
}

func (f FFmpeg) getSubtitleInputStr(subtitle subtitle.Subtitle) string {
	return fmt.Sprintf("./%s/%d.mkv", f.config.Folder.Temporary, subtitle.Id)
}

func (f FFmpeg) getSubtitleFilterStr(subtitle subtitle.Subtitle) string {
	return fmt.Sprintf(
		"subtitles=./%s/%d.srt:force_style='FontName=%s,Fontsize=%d',scale=%d:-1:flags=bicubic,fps=%d",
		f.config.Folder.Temporary,
		subtitle.Id,
		f.config.Gif.Subtitle.Font,
		f.config.Gif.Subtitle.Size,
		f.config.Gif.Resolution,
		f.config.Gif.Fps,
	)
}

func (f FFmpeg) getSubtitleOutputStr(subtitle subtitle.Subtitle) string {
	return fmt.Sprintf("./%s/%s/%s.gif", f.config.Folder.Output, f.config.Parameter.Episode, subtitle.Filename)
}
