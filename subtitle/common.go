package subtitle

import (
	"fmt"
	"oneliner-generator/filesystem"
	"time"
)

const timeFormat = "15:04:05.000"

func parseDuration(timeStrings []string, format string) (duration, error) {
	var duration duration

	var times []time.Time
	for _, t := range timeStrings {
		parsed, err := time.Parse(format, t)
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

func getFileName(id int, lines []string) string {
	var text string

	for _, line := range lines {
		text = fmt.Sprintf("%s %s", text, line)
	}

	return fmt.Sprintf("%d. %s", id, filesystem.SanitizeFileName(text))
}
