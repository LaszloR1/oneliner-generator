package config

type Config struct {
	Folder      folder `json:"folder"`
	Gif         gif    `json:"gif"`
	LengthCheck bool   `json:"length_check"`
	Parameter   parameter
}

type folder struct {
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

type parameter struct {
	Episode         string
	SkipCheckLength bool
}
