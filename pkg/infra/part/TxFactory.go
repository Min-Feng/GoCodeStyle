package part

import (
	"context"
	"database/sql/driver"

	"github.com/jmoiron/sqlx"
)

func NewTxFactory(db *sqlx.DB) *TxFactory {
	return &TxFactory{db: db}
}

type TxFactory struct {
	db *sqlx.DB
}

func (f *TxFactory) Tx() (driver.Tx, error) {
	return f.db.Beginx()
}

func (f *TxFactory) ContextWithTx(ctx context.Context) (context.Context, error) {
	tx, err := f.db.Beginx()
	ctx = context.WithValue(ctx, "tx", tx)
	return ctx, err
}
