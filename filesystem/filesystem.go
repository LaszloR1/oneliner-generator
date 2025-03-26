package filesystem

import (
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"strings"
)

const dirMode = 0755

func (f Filesystem) Setup() {
	os.RemoveAll(fmt.Sprintf("./%s", f.folder.Temporary))
	os.RemoveAll(fmt.Sprintf("./%s", f.folder.Output))

	for _, folder := range []string{f.folder.Input, f.folder.Output, f.folder.Temporary} {
		os.MkdirAll(fmt.Sprintf("./%s", folder), dirMode)
	}

	os.MkdirAll(fmt.Sprintf("./%s/%s", f.folder.Output, f.parameter.Episode), dirMode)
}

func (f Filesystem) cleanDir(folder string) {
	os.RemoveAll(fmt.Sprintf("./%s", folder))
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

	fmt.Println("Subtitles saved")

	return nil
}

func (f Filesystem) CreateTemp(filename string, content string) {
	os.WriteFile(fmt.Sprintf("./%s/%s", f.folder.Temporary, filename), []byte(content), os.ModeAppend)
}
