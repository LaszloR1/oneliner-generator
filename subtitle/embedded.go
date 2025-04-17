package subtitle

import (
	"oneliner-generator/config"
	"oneliner-generator/ffmpeg"
	"oneliner-generator/filesystem"
	"oneliner-generator/logger"
)

type embeddedParser struct {
	config config.Config
	fs     filesystem.Filesystem
	logger logger.Logger
	ffmpeg ffmpeg.FFmpeg
}

func NewEmbeddedParser(config config.Config, fs filesystem.Filesystem, logger logger.Logger, ffmpeg ffmpeg.FFmpeg) embeddedParser {
	return embeddedParser{
		config: config,
		fs:     fs,
		logger: logger,
		ffmpeg: ffmpeg,
	}
}

func (ep embeddedParser) Parse() ([]Subtitle, error) {
	ep.logger.Log(logger.Stage, "embedded subtitle parser")

	ep.ffmpeg.Extract()

	parser := NewSrtParser(ep.config, ep.fs, ep.logger)

	return parser.Parse()
}
