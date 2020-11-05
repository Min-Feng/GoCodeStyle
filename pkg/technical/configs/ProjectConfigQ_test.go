package configs_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"ddd/pkg/technical/configs"
	"ddd/pkg/technical/mock"
)

func TestProjectConfigRepo(t *testing.T) {
	suite.Run(t, new(ProjectConfigRepoTestSuite))
}

type ProjectConfigRepoTestSuite struct {
	suite.Suite
}

func (ts *ProjectConfigRepoTestSuite) TestFind() {
	tests := []struct {
		name               string
		repo               configs.ProjectConfigRepoQ
		expectedConfigName string
	}{
		{
			name:               "Read_Local_File_Dev",
			repo:               configs.NewLocalRepo("dev"),
			expectedConfigName: "dev",
		},
		{
			name:               "Read_Mock_Repo",
			repo:               mock.ProjectConfigRepo{},
			expectedConfigName: "mock",
		},
		// {
		// 	name:               "Read_Remote_Config",
		// 	repo:               configs.NewApolloRepo(""),
		// 	expectedConfigName: "apollo",
		// },
	}

	// 假設有讀到數值, 代表真的成功從相對應的位置 load config
	for _, tt := range tests {
		tt := tt
		ts.Run(tt.name, func() {
			actualConfig := tt.repo.QueryConfig()
			assert.Equal(ts.T(), tt.expectedConfigName, actualConfig.Name)
		})
	}
}
