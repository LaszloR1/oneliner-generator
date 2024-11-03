package main

import (
	"flag"
	"log"
	"oneliner-generator/caption"
	"oneliner-generator/util"
)

func main() {
	config := util.LoadConfig()

	fs := util.NewFileSystem(config)
	fs.SetupFolders()

	ep := flag.String("ep", "", "filename of the episode you want to parse")
	lc := flag.Bool("lc", true, "halts the program if the subtitle is visible for less than a frame (ffmpeg cannot deal with such clips)")

	flag.Parse()

	if *ep == "" {
		log.Fatal("You need to specify an episode with the -ep flag!")
	}

	if *lc == false {
		config.LengthCheck = false
	}

	fs.CreateFolderForEpisode(*ep)

	srt := caption.NewSrt(config, fs, *ep)
	subtitles := srt.Parse()
	srt.CreateTempSrts(subtitles)
	fs.SaveSubtitlesAsJson(*ep, subtitles)

	ffmpeg := caption.NewFFmpeg(config, subtitles, fs, *ep)
	ffmpeg.Run()
}
