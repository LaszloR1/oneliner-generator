package config

import (
	"encoding/json"
	"errors"
	"flag"
	"os"
	"strings"
)

func Parse() (Config, error) {
	config, err := parseConfig()
	if err != nil {
		return config, err
	}

	config.Parameter, err = parseParameter()
	if err != nil {
		return config, err
	}

	return config, nil
}

func parseConfig() (Config, error) {
	config := Config{}

	file, err := os.ReadFile("./config.json")
	if err != nil {
		return config, err
	}

	err = json.Unmarshal(file, &config)
	if err != nil {
		return config, err
	}

	return config, nil
}

func parseParameter() (Parameter, error) {
	file := flag.String("file", "", "file of the episode you want to parse")
	lc := flag.Bool("lc", true, "halts the program if the subtitle is visible for less than a frame (ffmpeg cannot deal with such clips)")

	flag.Parse()

	episode, format, err := parseInput(*file)
	if err != nil {
		return Parameter{}, err
	}

	return Parameter{
		Episode:         episode,
		Format:          format,
		SkipCheckLength: *lc,
	}, nil
}

func parseInput(file string) (string, string, error) {
	parts := strings.Split(file, ".")
	if len(parts) < 2 {
		return "", "", errors.New("Could not determine filetype from file.")
	}

	final := len(parts) - 1
	episode := strings.Join(parts[0:final], ".")
	format := parts[final]

	return episode, format, nil
}
