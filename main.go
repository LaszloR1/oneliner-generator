package main

import (
	"log"
	"oneliner-generator/caption"
	"oneliner-generator/config"
	"oneliner-generator/filesystem"
	"oneliner-generator/subtitle"
)

func main() {
	config, err := config.Parse()
	if err != nil {
		log.Fatal(err.Error())
	}

	fs := filesystem.New(config)
	fs.Setup()

	parser := subtitle.NewSubtitleParser(fs, config)
	subtitles, err := parser.Parse(config.Parameter.Episode)
	if err != nil {
		log.Fatal(err.Error())
	}

	err = fs.SavesAsJson(subtitles)
	if err != nil {
		log.Fatal(err.Error())
	}
	parser.CreateTempSubtitleSrts(subtitles)

	ffmpeg := caption.NewFFmpeg(config, fs)
	ffmpeg.Run(subtitles)
}
