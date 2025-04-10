package logger

import (
	"fmt"
	"slices"
)

func (l Logger) Log(mode Mode, message string) {
	if !l.isModeEnabled(mode) {
		return
	}

	fmt.Printf("[%s] %s\n", mode, message)
}

func (l Logger) isModeEnabled(mode Mode) bool {
	if mode == Critical {
		return true
	}

	return slices.Contains(l.config.Log.Types, string(mode))
}
