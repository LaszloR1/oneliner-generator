package subtitle

import (
	"bufio"
	"fmt"
	"oneliner-generator/config"
	"oneliner-generator/filesystem"
	"oneliner-generator/logger"
	"os"
	"strings"
)

type assParser struct {
	config config.Config
	fs     filesystem.Filesystem
	logger logger.Logger
}

func NewAssParser(config config.Config, fs filesystem.Filesystem, logger logger.Logger) assParser {
	return assParser{
		config: config,
		fs:     fs,
		logger: logger,
	}
}

const assSeparator = ","
const assTimeFormat = "15:04:05.00"

func (ap assParser) Parse() ([]Subtitle, error) {
	ap.logger.Log(logger.Stage, "ass subtitle parser")

	var subtitles []Subtitle

	file, err := os.Open(fmt.Sprintf("./%s/%s.ass", ap.config.Folder.Input, ap.config.Parameter.Episode))
	if err != nil {
		return subtitles, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	n := 1
	for scanner.Scan() {
		line := filesystem.DirtyBomFix(scanner.Text())
		if !strings.HasPrefix(line, "Dialogue:") {
			continue
		}

		parts := strings.Split(line, assSeparator)
		start := parts[1]
		end := parts[2]
		lines := strings.Split(strings.Join(parts[9:], assSeparator), "\\N")

		duration, err := parseDuration([]string{start, end}, assTimeFormat, ap.config.Parameter.SubtitleDelay)
		if err != nil {
			return subtitles, err
		}

		subtitles = append(subtitles, Subtitle{
			Id:       n,
			Lines:    lines,
			Duration: duration,
			Filename: getFileName(n, lines),
		})

		n++
	}

	return subtitles, nil
}
