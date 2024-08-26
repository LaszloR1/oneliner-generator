package types

type Config struct {
	InputFolder      string `env:"INPUT_FOLDER, required"`
	OutputFolder     string `env:"OUTPUT_FOLDER, required"`
	TempFolder       string `env:"TEMP_FOLDER, required"`
	SubtitleFontsize int    `env:"SUBTITLE_FONTSIZE, required"`
	GifResolution    int    `env:"GIF_RESOLUTION, required"`
	GifFramerate     int    `env:"GIF_FRAMERATE, required"`
}
