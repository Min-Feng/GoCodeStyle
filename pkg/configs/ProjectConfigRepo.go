package configs

import (
	"fmt"

	"ddd/pkg/helper/helperlog"
)

type ProjectConfigRepo interface {
	Find() *ProjectConfig
}

type ProjectConfig struct {
	Name       string          `configs:"name"`
	Port       string          `configs:"port"`
	AlarmEmail string          `configs:"alarm_email"`
	LogLevel   helperlog.Level `configs:"log_level"`
	MySQL      MySQL           `configs:"mysql"`
}

type MySQL struct {
	User        string `configs:"user"`
	Password    string `configs:"password"`
	Host        string `configs:"host"`
	Port        string `configs:"port"`
	Database    string `configs:"database"`
	MaxConn     int    `configs:"max_conn"`
	MaxIdleConn int    `configs:"max_idle_conn"`
}

func (c *MySQL) DSN() string {
	return fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?parseTime=true&loc=Local", c.User, c.Password, c.Host, c.Port, c.Database)
}
