package configs

import (
	"os"

	"ddd/gopkg2/domain"
)

func New() *domain.Config {
	os.Open("./config.yaml")
	// do something
	return new(domain.Config)
}
