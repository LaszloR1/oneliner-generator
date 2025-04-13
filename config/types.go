package config

type Config struct {
	Folder    Folder `json:"folder"`
	Gif       gif    `json:"gif"`
	Log       log    `json:"log"`
	Parameter Parameter
}

type Folder struct {
	Input     string `json:"input"`
	Output    string `json:"output"`
	Temporary string `json:"temporary"`
}

type gif struct {
	Fps        int      `json:"fps"`
	Resolution int      `json:"resolution"`
	Subtitle   subtitle `json:"subtitle"`
}

type subtitle struct {
	Font        string `json:"font"`
	Size        int    `json:"size"`
	CheckLength bool   `json:"check_length"`
}

type Parameter struct {
	Episode         string
	Format          string
	SkipCheckLength bool
}

type log struct {
	Types []string
}
