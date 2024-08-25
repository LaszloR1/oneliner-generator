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

	if len(os.Args) > 1 {
		ep = os.Args[1]
	}

	srt_file := path + ep + ".srt"
	mkv_file := path + ep + ".mkv"

	f, err := os.Open(srt_file)
	if err != nil {
		panic(err.Error())
	}

	caption.ClearTmp()
	subtitles := caption.ParseSrt(f)
	for i, s := range subtitles {
		fmt.Printf("%d/%d - %+v", i+1, len(subtitles), s)
		fmt.Printf("%+v\n", s)
		caption.TempSrt(s)
		ffmpeg.Trim(mkv_file, s)
		ffmpeg.AddSubtitles(s)
	}
}
