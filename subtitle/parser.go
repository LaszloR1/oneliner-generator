package subtitle

type Parser interface {
	Parse(string) ([]Subtitle, error)
}
