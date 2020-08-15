package configs

import (
	"ddd/pkg/loghelper"
	"ddd/pkg/repository/mysql"
)

type ProjectConfigRepo interface {
	Find() *ProjectConfig
}

type ProjectConfig struct {
	Name       string          `configs:"name"`
	Port       string          `configs:"port"`
	AlarmEmail string          `configs:"alarm_email"`
	LogLevel   loghelper.Level `configs:"log_level"`
	Mysql      mysql.Config    `configs:"mysql"`
}
