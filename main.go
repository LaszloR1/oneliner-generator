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
	if err := fs.Setup(); err != nil {
		log.Fatal(err.Error())
	}

	parser := subtitle.NewSubtitleParser(fs, config)
	subtitles, err := parser.Parse(config.Parameter.Episode)
	if err != nil {
		log.Fatal(err.Error())
	}

	ffmpeg := caption.NewFFmpeg(config, fs)
	ffmpeg.Run(subtitles)
}
