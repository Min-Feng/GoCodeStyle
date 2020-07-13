package configs

import (
	"os"
)

func New() *Config {
	os.Open("./config.yaml")
	// do something
	return new(Config)
}
