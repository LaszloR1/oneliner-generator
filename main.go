package main

import (
	"context"
	_ "github.com/joho/godotenv/autoload"
	"github.com/sethvargo/go-envconfig"
	"log"
	"oneliner-generator/caption"
	"oneliner-generator/types"
	"oneliner-generator/util"
	"os"
)

func main() {
	var config types.Config
	err := envconfig.Process(context.Background(), &config)
	if err != nil {
		log.Fatal(err)
	}

	fs := util.NewFileSystem(config)
	fs.SetupFolders()

	if len(os.Args) < 2 {
		log.Fatal("first argument not provided: it should be the name of the file without file extensions")
	}

	ep := os.Args[1]

	fs.CreateFolderForEpisode(ep)

	srt := caption.NewSrt(config, fs, ep)
	subtitles := srt.Parse()
	srt.CreateTempSrts(subtitles)

	ffmpeg := caption.NewFFmpeg(config, subtitles, fs, ep)
	ffmpeg.Run()
}
