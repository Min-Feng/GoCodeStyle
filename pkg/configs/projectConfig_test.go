package configs_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"ddd/pkg/configs"
	"ddd/pkg/infra/loghelper"
	"ddd/pkg/mock"
)

func init() {
	loghelper.InitAtUnitTest()
}

func TestProjectConfigStore(t *testing.T) {
	suite.Run(t, new(ProjectConfigStoreTestSuite))
}

type ProjectConfigStoreTestSuite struct {
	suite.Suite
}

func (ts *ProjectConfigStoreTestSuite) TestFind() {
	tests := []struct {
		name             string
		store            configs.ProjectConfigStore
		expectConfigName string
	}{
		{
			name:             "Read_Local_File_Dev",
			store:            configs.NewLocalProjectConfigStore("dev", "../../config"),
			expectConfigName: "dev",
		},
		{
			name:             "Read_Local_File_Test",
			store:            configs.NewLocalProjectConfigStore("test", "../../config"),
			expectConfigName: "test",
		},
		{
			name:             "Read_Local_File_Prod",
			store:            configs.NewLocalProjectConfigStore("prod", "../../config"),
			expectConfigName: "prod",
		},
		{
			name:             "Read_Mock_Store",
			store:            mock.ProjectConfigStore{},
			expectConfigName: "mock",
		},
		// {
		// 	name:             "Read_Remote_Config",
		// 	store:            configs.NewApolloProjectConfigStore(""),
		// 	expectConfigName: "apollo",
		// },
	}

	// 假設有讀到數值, 代表真的成功從相對應的位置 load config
	for _, tt := range tests {
		ts.Run(tt.name, func() {
			actualConfig := tt.store.Find()
			assert.Equal(ts.T(), tt.expectConfigName, actualConfig.Name)
		})
	}
}
