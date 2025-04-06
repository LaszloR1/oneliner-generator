package subtitle

import (
	"bufio"
	"errors"
	"fmt"
	"oneliner-generator/config"
	"oneliner-generator/filesystem"
	"os"
	"strconv"
	"strings"
	"time"
)

type LineType int

const (
	Empty LineType = iota
	Id
	Time
	Text
)

const timeFormat = "15:04:05.000"
const srtTimeFormat = "15:04:05,000"
const durationSeparator = " --> "

func (s Subtitle) getFullLine() string {
	return strings.Join(s.Lines, " ")
}

type srtParser struct {
	config config.Config
	fs     filesystem.Filesystem
}

func NewSrtParser(config config.Config, fs filesystem.Filesystem) srtParser {
	return srtParser{
		config: config,
		fs:     fs,
	}
}

func (sp srtParser) Parse(filename string) ([]Subtitle, error) {
	var subtitles []Subtitle

	file, err := os.Open(fmt.Sprintf("./%s/%s.srt", sp.config.Folder.Input, filename))
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
			duration, err := parseDuration(line)
			if err != nil {
				return subtitles, err
			}
			if sp.lessThanAFrame(duration.Length) {
				return subtitles, errors.New(fmt.Sprintf("Subtitle %d is less than a frame!", subtitle.Id))
			}

			subtitle.Duration = duration
		case Text:
			subtitle.Lines = append(subtitle.Lines, line)
		case Empty:
			subtitle.Filename = subtitle.getFileName()
			subtitles = append(subtitles, subtitle)
			subtitle = Subtitle{}
		}
	}

	err = sp.fs.SavesAsJson(subtitles)
	if err != nil {
		return subtitles, err
	}

	sp.createTempSubtitleSrts(subtitles)

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

func parseDuration(line string) (duration, error) {
	var duration duration

	timeStrings := strings.Split(line, " --> ")

	var times []time.Time
	for _, t := range timeStrings {
		parsed, err := time.Parse(srtTimeFormat, t)
		if err != nil {
			return duration, err
		}
		times = append(times, parsed)
	}

	duration.From = times[0]
	duration.To = times[1]
	duration.Length = duration.To.Sub(duration.From)

	return duration, nil
}

func (srt srtParser) lessThanAFrame(length time.Duration) bool {
	if srt.config.Parameter.SkipCheckLength {
		return false
	}

	if !srt.config.Gif.Subtitle.CheckLength {
		return false
	}

	if float64(1000)/float64(srt.config.Gif.Fps) < float64(length.Seconds()) {
		return false
	}

	return true

}

func (s Subtitle) getFileName() string {
	var text string

	for _, line := range s.Lines {
		text = fmt.Sprintf("%s %s", text, line)
	}

	return fmt.Sprintf("%d. %s", s.Id, filesystem.SanitizeFileName(text))
}

func (sp srtParser) createTempSubtitleSrts(subtitles []Subtitle) error {
	contents := []string{"1", "00:00:00,000 --> 00:01:00,000"}

	for _, subtitle := range subtitles {
		name := fmt.Sprintf("%d.srt", subtitle.Id)

		err := sp.fs.CreateTemp(name, strings.Join(append(contents, subtitle.Lines...), "\n"))
		if err != nil {
			return err
		}
	}

	return nil
}
