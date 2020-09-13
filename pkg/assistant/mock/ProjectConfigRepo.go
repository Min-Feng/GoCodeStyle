package mock

import (
	"ddd/pkg/assistant/configs"
)

type ProjectConfigRepo struct{}

func (ProjectConfigRepo) Find() *configs.ProjectConfig {
	return &configs.ProjectConfig{
		Name: "mock",
	}
}
