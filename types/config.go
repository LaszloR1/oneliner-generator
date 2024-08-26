package types

type Config struct {
	InputFolder  string `env:"INPUT_FOLDER, required"`
	OutputFolder string `env:"OUTPUT_FOLDER, required"`
	TempFolder   string `env:"TEMP_FOLDER, required"`
}
