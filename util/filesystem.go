package util

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/fs"
	"log"
	"oneliner-generator/types"
	"os"
	"regexp"
	"strings"
)

func LoadConfig() types.Config {
	var config types.Config

	file, err := os.ReadFile("./config.json")
	if err != nil {
		fmt.Println("Error: Couldn't load contents of the config file!")
		log.Fatal(err.Error())
	}

	err = json.Unmarshal(file, &config)
	if err != nil {
		fmt.Println(string(file))
		fmt.Println("Error: Couldn't parse the config file!")
		log.Fatal(err.Error())
	}

	return config
}

type FileSystem struct {
	config types.Config
}

func NewFileSystem(config types.Config) FileSystem {
	return FileSystem{config: config}
}

func (f FileSystem) SetupFolders() {
	f.ClearTmp()

	os.MkdirAll(fmt.Sprintf("./%s", f.config.InputFolder), fs.ModeDir)
	os.MkdirAll(fmt.Sprintf("./%s", f.config.OutputFolder), fs.ModeDir)
	os.MkdirAll(fmt.Sprintf("./%s", f.config.TempFolder), fs.ModeDir)
}

func (f FileSystem) CreateTempFile(name string, contents string) {
	file := fmt.Sprintf("./%s/%s", f.config.TempFolder, name)

	os.WriteFile(file, []byte(contents), os.ModeAppend)
}

func (f FileSystem) ClearTmp() {
	os.RemoveAll(fmt.Sprintf("./%s", f.config.TempFolder))
}

func (f FileSystem) CreateFolderForEpisode(name string) {
	os.MkdirAll(fmt.Sprintf("./%s/%s", f.config.OutputFolder, name), fs.ModeDir)
}

func (f FileSystem) DirtyBomFix(text string) string {
	return strings.ReplaceAll(text, "\ufeff", "")
}

func (f FileSystem) SanitizeFileName(text string) string {
	re1 := regexp.MustCompile("<[^>]*>")
	re2 := regexp.MustCompile(`[\\/:*?"<>|]`)

	safeName := re1.ReplaceAllString(text, "")
	safeName = re2.ReplaceAllString(safeName, "")
	safeName = strings.Trim(safeName, " .")

	maxLength := 255
	if len(safeName) > maxLength {
		safeName = safeName[:maxLength]
	}

	return safeName
}

func (f FileSystem) GetNextLine(scanner *bufio.Scanner) string {
	scanner.Scan()

	return scanner.Text()
}

func (f FileSystem) SaveSubtitlesAsJson(ep string, subtitles types.Subtitles) {
	jsonData, err := json.MarshalIndent(subtitles, "", "    ")
	if err != nil {
		log.Fatal(err.Error())
	}

	file := fmt.Sprintf("./%s/%s/subtitles.json\n", f.config.OutputFolder, ep)
	os.WriteFile(file, jsonData, os.ModeAppend)

	fmt.Printf("Subtitles saved to: %s", file)
}
