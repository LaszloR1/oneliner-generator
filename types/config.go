package types

type Config struct {
	InputFolder      string `json:"input_folder"`
	OutputFolder     string `json:"output_folder"`
	TempFolder       string `json:"temp_folder"`
	SubtitleFontsize int    `json:"subtitle_fontsize"`
	GifResolution    int    `json:"gif_resolution"`
	GifFramerate     int    `json:"gif_framerate"`
	LengthCheck      bool   `json:"length_check"`
}
