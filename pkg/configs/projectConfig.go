package configs

import (
	"ddd/pkg/infra/loghelper"
	"ddd/pkg/repository/mysql"
	"ddd/pkg/repository/redis"
)

type ProjectConfigStore interface {
	Find() *ProjectConfig
}

type ProjectConfig struct {
	Name       string          `configs:"name"`
	Port       string          `configs:"port"`
	AlarmEmail string          `configs:"alarm_email"`
	LogLevel   loghelper.Level `configs:"log_level"`

	Repo struct {
		MySQL struct {
			Writer mysql.Config `configs:"writer"`
			Reader mysql.Config `configs:"reader"`
		} `configs:"mysql"`

		Redis redis.Config `configs:"redis"`
	} `configs:"repo"`
}
