package unknown

import (
	"ddd/pkg/infra/loghelper"
	"ddd/pkg/repository/mysql"
	"ddd/pkg/repository/redis"
)

type ProjectConfig struct {
	Name       string          `config:"name"`
	Port       string          `config:"port"`
	AlarmEmail string          `config:"alarm_email"`
	LogLevel   loghelper.Level `config:"log_level"` // enum = [0, 1] = [Debug, Info]

	Repo struct {
		MySQL struct {
			Writer mysql.Config `config:"writer"`
			Reader mysql.Config `config:"reader"`
		} `config:"mysql"`

		Redis redis.Config `config:"redis"`
	} `config:"repo"`
}

type ProjectConfigStore interface {
	Find() *ProjectConfig
}
