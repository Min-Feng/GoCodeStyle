package configs_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"ddd/pkg/configs"
	"ddd/pkg/loghelper"
	"ddd/pkg/mock"
)

func init() {
	loghelper.ForUnitTest()
}

func TestProjectConfigRepo(t *testing.T) {
	suite.Run(t, new(ProjectConfigRepoTestSuite))
}

type ProjectConfigRepoTestSuite struct {
	suite.Suite
}

func (ts *ProjectConfigRepoTestSuite) TestFind() {
	loghelper.ForDeveloper()
	tests := []struct {
		name             string
		repo             configs.ProjectConfigRepo
		expectConfigName string
	}{
		{
			name:             "Read_Local_File_Dev",
			repo:             configs.NewLocalProjectConfigRepo("dev", "../../config"),
			expectConfigName: "dev",
		},
		// {
		// 	name:             "Read_Local_File_Test",
		// 	repo:            configs.NewLocalProjectConfigRepo("test", "../../config"),
		// 	expectConfigName: "test",
		// },
		// {
		// 	name:             "Read_Local_File_Prod",
		// 	repo:            configs.NewLocalProjectConfigRepo("prod", "../../config"),
		// 	expectConfigName: "prod",
		// },
		{
			name:             "Read_Mock_Store",
			repo:             mock.ProjectConfigRepo{},
			expectConfigName: "mock",
		},
		// {
		// 	name:             "Read_Remote_Config",
		// 	repo:            configs.NewApolloProjectConfigRepo(""),
		// 	expectConfigName: "apollo",
		// },
	}

	// 假設有讀到數值, 代表真的成功從相對應的位置 load config
	for _, tt := range tests {
		tt := tt
		ts.Run(tt.name, func() {
			actualConfig := tt.repo.Find()
			assert.Equal(ts.T(), tt.expectConfigName, actualConfig.Name)
		})
	}
}
