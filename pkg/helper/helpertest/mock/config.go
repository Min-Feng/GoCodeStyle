package mock

import (
	"ddd/pkg/configs"
)

func Config() *configs.ProjectConfig {
	return &configs.ProjectConfig{
		Name:       "",
		Port:       "",
		AlarmEmail: "",
		LogLevel:   "",
		MySQL: configs.MySQL{
			User:        "root",
			Password:    "1234",
			Host:        "127.0.0.1",
			Port:        "3306",
			Database:    "GoCodeStyle",
			MaxConn:     8,
			MaxIdleConn: 2,
		},
	}
}
