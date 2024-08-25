package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

func main() {
	path := "The_Wire/"
	ep := "S01E01"

	srt_file := path + ep + ".srt"
	mkv_file := path + ep + ".mkv"

	f, err := os.Open(srt_file)
	if err != nil {
		fmt.Println(err.Error())
		panic("no")
	}

	subtitles := parseSrt(f)
	for i, s := range subtitles {
		ffmpeg(mkv_file, srt_file, s)
		fmt.Printf("%d/%d - %s\n", i, len(subtitles), s.Subtitle)
	}
}

func ffmpeg(mkv_file string, srt_file string, s srt) {
	args := []string{
		"-i", mkv_file,
		//"-vf", fmt.Sprintf("subtitles=%s:force_style='Fontsize=24',scale=480:-1:flags=lanczos,fps=24", srt_file),
		"-vf", fmt.Sprintf("subtitles=%s:force_style='Fontsize=24',scale=480:-1:flags=bicubic,fps=24", srt_file),
		"-ss", s.Start,
		"-to", s.End,
		fmt.Sprintf("out/%d.%s.gif", s.Id, sanitize(s.Subtitle)),
	}

	cmd := exec.Command("ffmpeg", args...)
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}

func sanitize(name string) string {
	re := regexp.MustCompile(`[^\w\-. ]+`)
	safeName := re.ReplaceAllString(name, "")

	safeName = strings.Trim(safeName, " .")

	maxLength := 255
	if len(safeName) > maxLength {
		safeName = safeName[:maxLength]
	}

	return safeName
}
