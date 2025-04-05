package subtitle

import (
	"errors"
	"fmt"
	"oneliner-generator/config"
	"oneliner-generator/ffmpeg"
	"oneliner-generator/filesystem"
	"os"
)

func CreateParser(config config.Config, fs filesystem.Filesystem) (Parser, error) {
	if _, err := os.Stat(fmt.Sprintf("%s/%s.srt", config.Folder.Input, config.Parameter.Episode)); err == nil {
		return NewSrtParser(config, fs), nil
	}

	if _, err := os.Stat(fmt.Sprintf("%s/%s.mkv", config.Folder.Input, config.Parameter.Episode)); err == nil {
		return NewEmbeddedParser(config, fs, ffmpeg.New(config, fs)), nil
	}

	return nil, errors.New("Could not determine filetypes.")
}
