package logger

import "oneliner-generator/config"

type Mode string

const (
	Critical Mode = "critical"
	Stage    Mode = "stage"
	Render   Mode = "render"
)

type Logger struct {
	config config.Config
}

func NewLogger(config config.Config) Logger {
	return Logger{
		config: config,
	}
}
