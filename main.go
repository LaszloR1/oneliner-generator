package main

import (
	"log"
	"oneliner-generator/caption"
	"oneliner-generator/config"
	"oneliner-generator/util"
)

func main() {
	config, err := config.Parse()
	if err != nil {
		log.Fatal(err.Error())
	}

	fs := util.NewFileSystem(config)
	fs.SetupFolders()

	fs.CreateFolderForEpisode(config.Parameter.Episode)

	srt := caption.NewSrt(config, fs)
	subtitles := srt.Parse()
	srt.CreateTempSrts(subtitles)
	fs.SaveSubtitlesAsJson(config.Parameter.Episode, subtitles)

	ffmpeg := caption.NewFFmpeg(config, subtitles, fs, config.Parameter.Episode)
	ffmpeg.Run()
}
