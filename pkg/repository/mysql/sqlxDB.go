package mysql

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"

	"ddd/pkg/assistant/configs"
)

func NewDB(cfg *configs.MySQL) *sqlx.DB {
	db, err := sqlx.Connect("mysql", cfg.DSN())
	if err != nil {
		log.Error().Err(err).Msg("Connect mysql db failed:")
		return nil
	}

	db.SetMaxOpenConns(cfg.MaxConn)
	db.SetMaxIdleConns(cfg.MaxIdleConn)

	log.Info().Msg("Connect mysql db successfully")
	return db
}
