package subtitle

import (
	"time"
)

type Subtitle struct {
	Id       int
	Duration duration
	Lines    []string
	Filename string
}

type duration struct {
	From   time.Time
	To     time.Time
	Length time.Duration
}

type Parser interface {
	Parse() ([]Subtitle, error)
}
