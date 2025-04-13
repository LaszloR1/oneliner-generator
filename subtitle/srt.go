package subtitle

import (
	"bufio"
	"fmt"
	"oneliner-generator/config"
	"oneliner-generator/filesystem"
	"oneliner-generator/logger"
	"os"
	"strconv"
	"strings"
)

type LineType int

const (
	Empty LineType = iota
	Id
	Time
	Text
)

const srtTimeFormat = "15:04:05,000"
const durationSeparator = " --> "

func (s Subtitle) getFullLine() string {
	return strings.Join(s.Lines, " ")
}

type srtParser struct {
	config config.Config
	fs     filesystem.Filesystem
	logger logger.Logger
}

func NewSrtParser(config config.Config, fs filesystem.Filesystem, logger logger.Logger) srtParser {
	return srtParser{
		config: config,
		fs:     fs,
		logger: logger,
	}
}

func (sp srtParser) Parse() ([]Subtitle, error) {
	sp.logger.Log(logger.Stage, "srt subtitle parser")

	var subtitles []Subtitle

	file, err := os.Open(fmt.Sprintf("./%s/%s.srt", sp.config.Folder.Input, sp.config.Parameter.Episode))
	if err != nil {
		return subtitles, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	subtitle := Subtitle{}
	for scanner.Scan() {
		line := filesystem.DirtyBomFix(scanner.Text())

		switch identifyLine(line) {
		case Id:
			id, err := strconv.Atoi(line)
			if err != nil {
				return subtitles, err
			}
			subtitle.Id = id
		case Time:
			duration, err := parseDuration(strings.Split(line, " --> "), srtTimeFormat)
			if err != nil {
				return subtitles, err
			}

			subtitle.Duration = duration
		case Text:
			subtitle.Lines = append(subtitle.Lines, line)
		case Empty:
			subtitle.Filename = getFileName(subtitle.Id, subtitle.Lines)
			subtitles = append(subtitles, subtitle)
			subtitle = Subtitle{}
		}
	}

	return subtitles, nil
}

func identifyLine(line string) LineType {
	line = strings.TrimSpace(line)

	if line == "" {
		return Empty
	}

	if _, err := strconv.Atoi(line); err == nil {
		return Id
	}

	if strings.Contains(line, "-->") {
		return Time
	}

	return Text
}
