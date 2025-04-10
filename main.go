package main

import (
	"log"
	"oneliner-generator/config"
	"oneliner-generator/ffmpeg"
	"oneliner-generator/filesystem"
	logger_module "oneliner-generator/logger"
	"oneliner-generator/subtitle"
)

func main() {
	config, err := config.Parse()
	if err != nil {
		log.Fatal(err.Error())
	}

	logger := logger_module.NewLogger(config)

	fs := filesystem.New(config, logger)
	if err := fs.Setup(); err != nil {
		log.Fatal(err.Error())
	}

	parser, err := subtitle.CreateParser(config, fs, logger)
	if err != nil {
		log.Fatal(err.Error())
	}

	subtitles, err := parser.Parse(config.Parameter.Episode)
	if err != nil {
		log.Fatal(err.Error())
	}

	generator := subtitle.NewGenerator(config, fs, logger, ffmpeg.New(config, fs, logger))
	if err := generator.Run(subtitles); err != nil {
		log.Fatal(err.Error())
	}

	logger.Log(logger_module.Stage, "done")
}
