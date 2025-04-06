package subtitle

import (
	"errors"
	"fmt"
)

func lengthCheck(subtitles []Subtitle, fps int) error {
	for _, subtitle := range subtitles {
		if lessThanAFrame(subtitle.Duration.Length.Seconds(), fps) {
			return errors.New(fmt.Sprintf("Subtitle with Id %d is less than a frame long!", subtitle.Id))
		}
	}

	return nil
}

func lessThanAFrame(length float64, fps int) bool {
	return float64(length) < float64(1000)/float64(fps)
}
