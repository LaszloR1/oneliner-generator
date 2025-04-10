package filesystem

import (
	"encoding/json"
	"fmt"
	"oneliner-generator/logger"
	"os"
	"regexp"
	"strings"
)

const dirMode = 0755

func (f Filesystem) Setup() error {
	f.logger.Log(logger.Stage, "directory setup")

	for _, folder := range []string{f.folder.Temporary, f.folder.Output} {
		if err := os.RemoveAll(fmt.Sprintf("./%s", folder)); err != nil {
			return err
		}
	}

	for _, folder := range []string{f.folder.Input, f.folder.Output, f.folder.Temporary} {
		if err := os.MkdirAll(fmt.Sprintf("./%s", folder), dirMode); err != nil {
			return err
		}
	}

	if err := os.MkdirAll(fmt.Sprintf("./%s/%s", f.folder.Output, f.parameter.Episode), dirMode); err != nil {
		return err
	}

	return nil
}

func DirtyBomFix(text string) string {
	return strings.ReplaceAll(text, "\ufeff", "")
}

func SanitizeFileName(text string) string {
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

func (f Filesystem) SavesAsJson(subtitles any) error {
	jsonData, err := json.MarshalIndent(subtitles, "", "    ")
	if err != nil {
		return err
	}

	if err := os.WriteFile(fmt.Sprintf("./%s/%s/subtitles.json", f.folder.Output, f.parameter.Episode), jsonData, os.ModeAppend); err != nil {
		return err
	}

	f.logger.Log(logger.Stage, "json export")

	return nil
}

func (f Filesystem) CreateTemp(filename string, content string) error {
	return os.WriteFile(fmt.Sprintf("./%s/%s", f.folder.Temporary, filename), []byte(content), os.ModeAppend)
}
