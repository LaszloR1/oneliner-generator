package types

type Subtitle struct {
	Id       int
	From     string
	To       string
	Line1    string
	Line2    string
	Filename string
	Duration string
}

type Subtitles []Subtitle
