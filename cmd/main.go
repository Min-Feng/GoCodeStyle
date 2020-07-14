package main

import (
	"fmt"

	"ddd/pkg/repository/local"
	"ddd/pkg/repository/remote"

	"ddd/pkg/domain"
)

func main() {
	var c domain.Config
	c = NewConfig(LOCAL)

	fmt.Println(c)
}

type ConfigType string

const LOCAL ConfigType = "local"
const REMOTE ConfigType = "remote"

func NewConfig(c ConfigType) domain.Config {
	var store domain.ConfigStore
	switch c {
	case LOCAL:
		store = local.NewConfigStore()
	case REMOTE:
		store = remote.NewConfigStore()
	}
	config, _ := store.Find()
	return config
}
