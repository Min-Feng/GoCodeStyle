package mock

import (
	"ddd/pkg/technical/configs"
)

type ProjectConfigRepo struct{}

func (ProjectConfigRepo) Find() *configs.ProjectConfig {
	return &configs.ProjectConfig{
		Name: "mock",
	}
}
