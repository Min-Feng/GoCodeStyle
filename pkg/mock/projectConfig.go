package mock

import (
	"ddd/pkg/configs"
)

type ProjectConfigStore struct{}

func (ProjectConfigStore) Find() *configs.ProjectConfig {
	return &configs.ProjectConfig{
		Name: "mock",
	}
}

func NewProjectConfig() *configs.ProjectConfig {
	return &configs.ProjectConfig{
		Name: "mock",
	}
}
