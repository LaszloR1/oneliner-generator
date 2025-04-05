package main

import (
	"fmt"
	"log"
	"oneliner-generator/config"
	"oneliner-generator/ffmpeg"
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

	parser, err := subtitle.CreateParser(config, fs)
	if err != nil {
		log.Fatal(err.Error())
	}

	subtitles, err := parser.Parse(config.Parameter.Episode)
	if err != nil {
		log.Fatal(err.Error())
	}

	generator := subtitle.NewGenerator(ffmpeg.New(config, fs))
	if err := generator.Run(subtitles); err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("All done!")
}
