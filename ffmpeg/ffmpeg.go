package ffmpeg

import (
	"oneliner-generator/logger"
	"os/exec"
)

func (f FFmpeg) execute(args []string) error {
	cmd := exec.Command("ffmpeg", args...)
	if out, err := cmd.CombinedOutput(); err != nil {
		f.logger.Log(logger.Critical, string(out))

		return err
	}

	return nil
}
