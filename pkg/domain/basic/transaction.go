package basic

import (
	"context"
	"database/sql/driver"
)

type TxFactory interface {
	Tx() (driver.Tx, error)

	// 不希望函數簽名太醜
	// 所以用 context.WithValue 將 driver.Tx 包在 ctx
	ContextWithTx(ctx context.Context) (context.Context, error)
}

// HandleErrorByRollback
// 參數 err 代表外部錯誤, 預期不等於 nil
// 回傳 nil 代表 rollback 成功
// 依據業務 有不同的錯誤處理方式
// rollback 成功不代表 已經處理 外部錯誤
func HandleErrorByRollback(err error, tx driver.Tx) error {
	if err != nil {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			return rollbackErr
		}
	}
	return nil
}
