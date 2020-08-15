package mysql

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
)

func NewDB(cfg *Config) *sqlx.DB {
	db, err := sqlx.Connect("mysql", cfg.DSN())
	if err != nil {
		log.Fatal().Err(err).Msg("Connect mysql db failed:")
		return nil
	}

	db.SetMaxOpenConns(cfg.MaxConn)
	db.SetMaxIdleConns(cfg.MaxIdleConn)

	log.Info().Msg("Connect mysql db successfully")
	return db
}

type Config struct {
	User        string `configs:"user"`
	Password    string `configs:"password"`
	Host        string `configs:"host"`
	Port        string `configs:"port"`
	Database    string `configs:"database"`
	MaxConn     int    `configs:"max_conn"`
	MaxIdleConn int    `configs:"max_idle_conn"`
}

func (c *Config) DSN() string {
	return fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?parseTime=true&loc=Local", c.User, c.Password, c.Host, c.Port, c.Database)
}
