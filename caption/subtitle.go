package caption

import (
	"fmt"
	"regexp"
	"strings"
	"time"
)

type Subtitle struct {
	Id       int
	From     string
	To       string
	Line1    string
	Line2    string
	Filename string
	Duration string
}

type Subtitles []Subtitle

func (s Subtitle) generateFilename() Subtitle {
	text := s.Line1 + " " + s.Line2
	re := regexp.MustCompile(`[^\w\-. ]+`)
	safeName := re.ReplaceAllString(text, "")

	safeName = strings.Trim(safeName, " .")

	maxLength := 255
	if len(safeName) > maxLength {
		safeName = safeName[:maxLength]
	}

	s.Filename = safeName

	return s
}

func (s Subtitle) generateDuration() Subtitle {
	const layout = "15:04:05.000"
	t1, _ := time.Parse(layout, s.From)
	t2, _ := time.Parse(layout, s.To)

	duration := t2.Sub(t1)
	s.Duration = fmt.Sprintf("%.3f", duration.Seconds())

	return s
}
