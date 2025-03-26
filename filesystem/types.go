package filesystem

import "oneliner-generator/config"

type Filesystem struct {
	folder    config.Folder
	parameter config.Parameter
}

func New(config config.Config) Filesystem {
	return Filesystem{
		folder:    config.Folder,
		parameter: config.Parameter,
	}
}
