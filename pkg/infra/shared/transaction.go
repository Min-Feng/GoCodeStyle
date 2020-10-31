package shared

import (
	"context"

	"github.com/jmoiron/sqlx"
)

// GetTxOrDB is used to support that part.TxFactory's method ContextWithTx
func GetTxOrDB(ctx context.Context, db *sqlx.DB) (ext sqlx.ExtContext) {
	if ctx == nil {
		return db
	}

	externalTx, ok := ctx.Value("tx").(*sqlx.Tx)
	if ok {
		return externalTx
	}
	return db
}
