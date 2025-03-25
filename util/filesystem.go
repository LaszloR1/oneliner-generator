package util

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/fs"
	"log"
	"oneliner-generator/config"
	"oneliner-generator/types"
	"os"
	"regexp"
	"strings"
)

type FileSystem struct {
	config config.Config
}

func NewFileSystem(config config.Config) FileSystem {
	return FileSystem{config: config}
}

func (f FileSystem) SetupFolders() {
	f.ClearTmp()

	os.MkdirAll(fmt.Sprintf("./%s", f.config.Folder.Input), fs.ModeDir)
	os.MkdirAll(fmt.Sprintf("./%s", f.config.Folder.Output), fs.ModeDir)
	os.MkdirAll(fmt.Sprintf("./%s", f.config.Folder.Temporary), fs.ModeDir)
}

func (f FileSystem) CreateTempFile(name string, contents string) {
	file := fmt.Sprintf("./%s/%s", f.config.Folder.Temporary, name)

	os.WriteFile(file, []byte(contents), os.ModeAppend)
}

func (f FileSystem) ClearTmp() {
	os.RemoveAll(fmt.Sprintf("./%s", f.config.Folder.Temporary))
}

func (f FileSystem) CreateFolderForEpisode(name string) {
	os.MkdirAll(fmt.Sprintf("./%s/%s", f.config.Folder.Output, name), fs.ModeDir)
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

	file := fmt.Sprintf("./%s/%s/subtitles.json\n", f.config.Folder.Output, ep)
	os.WriteFile(file, jsonData, os.ModeAppend)

	fmt.Printf("Subtitles saved to: %s", file)
}
