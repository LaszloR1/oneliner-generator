package filesystem

import (
	"oneliner-generator/config"
	"oneliner-generator/logger"
)

type Filesystem struct {
	folder    config.Folder
	parameter config.Parameter
	logger    logger.Logger
}

func New(config config.Config, logger logger.Logger) Filesystem {
	return Filesystem{
		folder:    config.Folder,
		parameter: config.Parameter,
		logger:    logger,
	}
}
