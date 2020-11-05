package mock

import (
	"ddd/pkg/technical/configs"
)

type ProjectConfigRepo struct{}

func (ProjectConfigRepo) QueryConfig() *configs.ProjectConfig {
	return &configs.ProjectConfig{
		Name: "mock",
	}
}
