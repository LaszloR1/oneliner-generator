package ffmpeg

import (
	"fmt"
	"os/exec"
)

func (f FFmpeg) execute(args []string) error {
	cmd := exec.Command("ffmpeg", args...)
	if out, err := cmd.CombinedOutput(); err != nil {
		fmt.Println(string(out))
		return err
	}

	return nil
}
