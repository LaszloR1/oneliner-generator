package caption

import (
	"bufio"
	"fmt"
	"log"
	"oneliner-generator/types"
	"oneliner-generator/util"
	"os"
	"strconv"
	"strings"
	"time"
)

type Srt struct {
	config     types.Config
	filesystem util.FileSystem
	name       string
}

func NewSrt(config types.Config, filesystem util.FileSystem, name string) Srt {
	return Srt{
		config:     config,
		filesystem: filesystem,
		name:       name,
	}
}

func (s Srt) Parse() types.Subtitles {
	f, err := os.Open(fmt.Sprintf("%s/%s.srt", s.config.InputFolder, s.name))
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	subtitles := s.parseLines(f)

	return subtitles
}

func (s Srt) CreateTempSrts(subtitles types.Subtitles) {
	for _, subtitle := range subtitles {
		name := fmt.Sprintf("%d.srt", subtitle.Id)
		contents := fmt.Sprintf("1\n00:00:00,000 --> 00:01:00,000\n%s\n%s\n", subtitle.Line1, subtitle.Line2)

		s.filesystem.CreateTempFile(name, contents)
	}
}

func (s Srt) parseLines(f *os.File) types.Subtitles {
	var subtitles types.Subtitles

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		id, err := strconv.Atoi(s.filesystem.DirtyBomFix(scanner.Text()))
		if err != nil {
			continue
		}

		start, end := s.parseTimeSpan(scanner)
		l1, l2 := s.parseSubtitles(scanner)
		duration := s.generateDuration(start, end)
		subtitle := types.Subtitle{
			Id:       id,
			From:     start,
			To:       end,
			Line1:    l1,
			Line2:    l2,
			Duration: fmt.Sprintf("%.3f", duration),
			Filename: s.generateFileName(id, l1, l2),
		}

		if duration < 1.0/float64(s.config.GifFramerate) {
			fmt.Printf("Error: The duration of the following subtitle is less than a frame: %.3fs. Cannot proceed!\n", duration)
			log.Fatal(fmt.Sprintf("%+v\n", subtitle))
		}

		subtitles = append(subtitles, subtitle)
	}

	return subtitles
}

func (s Srt) parseTimeSpan(scanner *bufio.Scanner) (string, string) {
	arr := strings.Split(s.filesystem.GetNextLine(scanner), " --> ")

	return strings.ReplaceAll(arr[0], ",", "."), strings.ReplaceAll(arr[1], ",", ".")
}

func (s Srt) parseSubtitles(scanner *bufio.Scanner) (string, string) {
	l1 := s.filesystem.GetNextLine(scanner)
	l2 := s.filesystem.GetNextLine(scanner)

	return l1, l2
}

func (s Srt) generateDuration(from string, to string) float64 {
	const layout = "15:04:05.000"

	t1, _ := time.Parse(layout, from)
	t2, _ := time.Parse(layout, to)

	duration := t2.Sub(t1)

	return duration.Seconds()
}

func (s Srt) generateFileName(id int, l1 string, l2 string) string {
	text := l1

	if l2 != "" {
		text = text + " " + l2
	}

	text = fmt.Sprintf("%d. %s.gif", id, s.filesystem.SanitizeFileName(text))

	return text
}
