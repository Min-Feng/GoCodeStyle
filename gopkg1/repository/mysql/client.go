package mysql

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	"ddd/gopkg1/infra/configs"
)

type Client struct {
	DB *sql.DB
}

func NewClient(cfg *configs.Config) *Client {
	dsn := fmt.Sprintf("%v:%v@/dbnam", cfg.MySQLUser, cfg.MySQLPassword)
	db, _ := sql.Open("mysql", dsn)
	return &Client{db}
}
