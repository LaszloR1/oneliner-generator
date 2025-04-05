package subtitle

import (
	"fmt"
	"oneliner-generator/ffmpeg"
)

type Generator struct {
	ffmpeg ffmpeg.FFmpeg
}

func NewGenerator(ffmpeg ffmpeg.FFmpeg) Generator {
	return Generator{
		ffmpeg: ffmpeg,
	}
}

func (g Generator) Run(subtitles []Subtitle) error {
	for i, subtitle := range subtitles {
		fmt.Printf("%d/%d - %+v\n", i+1, len(subtitles), subtitle)

		if err := g.ffmpeg.Trim(subtitle.Id, subtitle.Duration.From, subtitle.Duration.Length); err != nil {
			return err
		}

		if err := g.ffmpeg.AddSubtitles(subtitle.Id, subtitle.Filename); err != nil {
			return err
		}
	}

	return nil
}
