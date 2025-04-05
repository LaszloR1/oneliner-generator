package subtitle

import (
	"fmt"
	"oneliner-generator/config"
	"oneliner-generator/ffmpeg"
	"oneliner-generator/filesystem"
)

type embeddedParser struct {
	config config.Config
	fs     filesystem.Filesystem
	ffmpeg ffmpeg.FFmpeg
}

func NewEmbeddedParser(config config.Config, fs filesystem.Filesystem, ffmpeg ffmpeg.FFmpeg) embeddedParser {
	return embeddedParser{
		config: config,
		fs:     fs,
		ffmpeg: ffmpeg,
	}
}

func (ep embeddedParser) Parse(filename string) ([]Subtitle, error) {
	fmt.Println("Extracting embedded subtitles...")

	ep.ffmpeg.Extract()

	fmt.Println("Extracted embedded subtitles")

	parser := NewSrtParser(ep.config, ep.fs)

	return parser.Parse(filename)
}
