package config

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
)

func Parse() (Config, error) {
	config, err := parseConfig()
	if err != nil {
		return config, err
	}

	config.Parameter = parseParameter()

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

func parseParameter() parameter {
	ep := flag.String("ep", "", "filename of the episode you want to parse")
	lc := flag.Bool("lc", true, "halts the program if the subtitle is visible for less than a frame (ffmpeg cannot deal with such clips)")

	flag.Parse()

	fmt.Println(*ep)

	return parameter{
		Episode:         *ep,
		SkipCheckLength: *lc,
	}
}
