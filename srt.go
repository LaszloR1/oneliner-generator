package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type srt struct {
	Id       int
	Start    string
	End      string
	Subtitle string
}

func parseSrt(f *os.File) []srt {
	defer f.Close()

	var subtitles []srt

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		var start, end, subtitle string

		id, err := strconv.Atoi(scanner.Text())
		if err != nil {
			continue
		}

		start, end = parseSpan(scanner)
		subtitle = parseSubtitles(scanner)
		subtitles = append(subtitles, srt{Id: id, Start: start, End: end, Subtitle: subtitle})

	}

	return subtitles
}

func parseSpan(scanner *bufio.Scanner) (string, string) {
	scanner.Scan()
	arr := strings.Split(scanner.Text(), " --> ")
	return strings.ReplaceAll(arr[0], ",", "."), strings.ReplaceAll(arr[1], ",", ".")
}

func parseSubtitles(scanner *bufio.Scanner) string {
	var subtitle string
	scanner.Scan()
	subtitle = scanner.Text()
	scanner.Scan()
	subtitle = subtitle + " " + scanner.Text()

	subtitle = strings.ReplaceAll(subtitle, `"`, `''`)

	return subtitle
}
