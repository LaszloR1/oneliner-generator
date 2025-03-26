package subtitle

import (
	"oneliner-generator/config"
	"oneliner-generator/filesystem"
	"time"
)

type Subtitle struct {
	Id       int
	Duration duration
	Lines    []string
	Filename string
}

type duration struct {
	From   time.Time
	To     time.Time
	Length time.Duration
}

type SubtitleParser struct {
	fs     filesystem.Filesystem
	config config.Config
}

func NewSubtitleParser(fs filesystem.Filesystem, config config.Config) SubtitleParser {
	return SubtitleParser{
		fs:     fs,
		config: config,
	}
}
