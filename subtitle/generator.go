package subtitle

import (
	"fmt"
	"oneliner-generator/config"
	"oneliner-generator/ffmpeg"
	"oneliner-generator/filesystem"
	"oneliner-generator/logger"
	"strings"
)

type Generator struct {
	config config.Config
	fs     filesystem.Filesystem
	logger logger.Logger
	ffmpeg ffmpeg.FFmpeg
}

func NewGenerator(config config.Config, fs filesystem.Filesystem, logger logger.Logger, ffmpeg ffmpeg.FFmpeg) Generator {
	return Generator{
		config: config,
		fs:     fs,
		logger: logger,
		ffmpeg: ffmpeg,
	}
}

func (g Generator) Run(subtitles []Subtitle) error {
	if err := g.validate(subtitles); err != nil {
		return err
	}

	if err := g.fs.SavesAsJson(subtitles); err != nil {
		return err
	}

	if err := g.createTempSubtitleSrts(subtitles); err != nil {
		return err
	}

	return g.generate(subtitles)
}

func (g Generator) generate(subtitles []Subtitle) error {
	g.logger.Log(logger.Stage, "gif generator")

	for i, subtitle := range subtitles {
		g.logger.Log(logger.Render, fmt.Sprintf("%d/%d - %+v\n", i+1, len(subtitles), subtitle))

		if err := g.ffmpeg.Trim(subtitle.Id, subtitle.Duration.From, subtitle.Duration.Length); err != nil {
			return err
		}

		if err := g.ffmpeg.AddSubtitles(subtitle.Id, subtitle.Filename); err != nil {
			return err
		}
	}

	return nil
}

func (g Generator) validate(subtitles []Subtitle) error {
	if g.config.Gif.Subtitle.CheckLength && !g.config.Parameter.SkipCheckLength {
		if err := lengthCheck(subtitles, g.config.Gif.Fps); err != nil {
			return err
		}
	}

	return nil
}
func (g Generator) createTempSubtitleSrts(subtitles []Subtitle) error {
	contents := []string{"1", "00:00:00,000 --> 00:01:00,000"}

	for _, subtitle := range subtitles {
		name := fmt.Sprintf("%d.srt", subtitle.Id)

		err := g.fs.CreateTemp(name, strings.Join(append(contents, subtitle.Lines...), "\n"))
		if err != nil {
			return err
		}
	}

	return nil
}
