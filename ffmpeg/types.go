package ffmpeg

import (
	"oneliner-generator/config"
	"oneliner-generator/filesystem"
	"oneliner-generator/logger"
)

type FFmpeg struct {
	config config.Config
	fs     filesystem.Filesystem
	logger logger.Logger
}

func New(config config.Config, fs filesystem.Filesystem, logger logger.Logger) FFmpeg {
	return FFmpeg{
		config: config,
		fs:     fs,
		logger: logger,
	}
}
