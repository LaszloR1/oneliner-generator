package subtitle

import (
	"errors"
	"fmt"
	"oneliner-generator/config"
	"oneliner-generator/ffmpeg"
	"oneliner-generator/filesystem"
	"oneliner-generator/logger"
	"os"
)

func CreateParser(config config.Config, fs filesystem.Filesystem, logger logger.Logger) (Parser, error) {
	if _, err := os.Stat(fmt.Sprintf("%s/%s.ass", config.Folder.Input, config.Parameter.Episode)); err == nil {
		return NewAssParser(config, fs, logger), nil
	}

	if _, err := os.Stat(fmt.Sprintf("%s/%s.srt", config.Folder.Input, config.Parameter.Episode)); err == nil {
		return NewSrtParser(config, fs, logger), nil
	}

	if _, err := os.Stat(fmt.Sprintf("%s/%s.mkv", config.Folder.Input, config.Parameter.Episode)); err == nil {
		return NewEmbeddedParser(config, fs, logger, ffmpeg.New(config, fs, logger)), nil
	}

	return nil, errors.New("Could not determine filetypes.")
}
