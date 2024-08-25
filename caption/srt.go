package caption

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ParseSrt(f *os.File) Subtitles {
	defer f.Close()

	var subtitles Subtitles

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		id, err := strconv.Atoi(scanner.Text())
		if err != nil {
			continue
		}

		start, end := parseSpan(scanner)
		l1, l2 := parseSubtitles(scanner)
		subtitles = append(subtitles, Subtitle{Id: id, From: start, To: end, Line1: l1, Line2: l2}.generateFilename().generateDuration())
	}

	return subtitles
}

func parseSpan(scanner *bufio.Scanner) (string, string) {
	scanner.Scan()
	arr := strings.Split(scanner.Text(), " --> ")
	return strings.ReplaceAll(arr[0], ",", "."), strings.ReplaceAll(arr[1], ",", ".")
}

func parseSubtitles(scanner *bufio.Scanner) (string, string) {
	scanner.Scan()
	l1 := scanner.Text()
	scanner.Scan()
	l2 := scanner.Text()

	return l1, l2
}

func TempSrt(s Subtitle) {
	os.WriteFile(fmt.Sprintf("_/tmp/%d.srt", s.Id), []byte(fmt.Sprintf("1\n00:00:00,000 --> 00:01:00,000\n%s\n%s\n", s.Line1, s.Line2)), os.ModeAppend)
}
