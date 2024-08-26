package util

import (
	"bufio"
	"fmt"
	"io/fs"
	"oneliner-generator/types"
	"os"
	"regexp"
	"strings"
)

type FileSystem struct {
	config types.Config
}

func NewFileSystem(config types.Config) FileSystem {
	return FileSystem{config: config}
}

func (f FileSystem) SetupFolders() {
	f.ClearTmp()

	os.MkdirAll(f.config.InputFolder, fs.ModeDir)
	os.MkdirAll(f.config.OutputFolder, fs.ModeDir)
	os.MkdirAll(f.config.TempFolder, fs.ModeDir)
}

func (f FileSystem) CreateTempFile(name string, contents string) {
	file := fmt.Sprintf("./%s/%s", f.config.TempFolder, name)

	os.WriteFile(file, []byte(contents), os.ModeAppend)
}

func (f FileSystem) ClearTmp() {
	os.RemoveAll(f.config.TempFolder)
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