package util

import (
	"fmt"
	"io/fs"
	"os"
	"regexp"
	"strings"
)

func TmpFile(name string, contents string) {
	os.WriteFile(fmt.Sprintf("_/tmp/%s", name), []byte(contents), os.ModeAppend)
}

func ClearTmp() {
	os.RemoveAll("_/tmp/")
	os.MkdirAll("_/tmp/", fs.ModeDir)
}

func DirtyBomFix(text string) string {
	return strings.ReplaceAll(text, "\ufeff", "")
}

func SanitizeFileName(text string) string {
	re1 := regexp.MustCompile("<[^>]*>")
	re2 := regexp.MustCompile(`[\\/:*?"<>|]`) // Include `()` to keep parentheses

	safeName := re1.ReplaceAllString(text, "")
	safeName = re2.ReplaceAllString(safeName, "")
	safeName = strings.Trim(safeName, " .")

	maxLength := 255
	if len(safeName) > maxLength {
		safeName = safeName[:maxLength]
	}

	return safeName
}
