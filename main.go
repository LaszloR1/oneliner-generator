package main

import (
	"fmt"
	"oneliner-generator/caption"
	"oneliner-generator/ffmpeg"
	"os"
)

func main() {
	path := "_/The_Wire/"
	ep := "S01E01"

	srt_file := path + ep + ".srt"
	mkv_file := path + ep + ".mkv"

	f, err := os.Open(srt_file)
	if err != nil {
		fmt.Println(err.Error())
		panic("no")
	}

	subtitles := caption.ParseSrt(f)
	for _, s := range subtitles {
		fmt.Printf("%+v\n", s)
		caption.TempSrt(s)
		ffmpeg.Trim(mkv_file, s)
		ffmpeg.AddSubtitles(s)
		//ffmpeg(mkv_file, srt_file, s)
		//fmt.Printf("%d/%d - %s\n", i, len(subtitles), s.Subtitle)
	}
}
