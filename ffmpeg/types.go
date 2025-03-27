package ffmpeg

import (
	"oneliner-generator/config"
	"oneliner-generator/filesystem"
)

type FFmpeg struct {
	config config.Config
	fs     filesystem.Filesystem
}

func New(config config.Config, fs filesystem.Filesystem) FFmpeg {
	return FFmpeg{
		config: config,
		fs:     fs,
	}
}
