package mysql

import (
	"database/sql"
	"fmt"

	"ddd/gopkg2/domain"

	_ "github.com/go-sql-driver/mysql"
)

type Client struct {
	DB *sql.DB
}

func NewClient(cfg *domain.Config) *Client {
	dsn := fmt.Sprintf("%v:%v@/dbnam", cfg.MySQLUser, cfg.MySQLPassword)
	db, _ := sql.Open("mysql", dsn)
	return &Client{db}
}
