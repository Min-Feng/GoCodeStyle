package mock

import (
	"ddd/pkg/configs"
)

type ProjectConfigRepo struct{}

func (ProjectConfigRepo) Find() *configs.ProjectConfig {
	return &configs.ProjectConfig{
		Name: "mock",
	}
}
