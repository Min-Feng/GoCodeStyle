package uow

import (
	"context"
	"database/sql/driver"

	"github.com/jmoiron/sqlx"
)

// TxFactory
// 用來控制聚合之間的資料強一致性, 是否要在同一個 tx
// 若沒傳入 tx 到 repo
// 則該聚合的資料強一致性 由該聚合自己負責
// 與其他的聚合不會同步
//
// 不希望 repo 函數簽名太醜
// 所以用 context.WithValue 將 driver.Tx 包在 ctx
//
// 日後有機會遭遇 使用人數大幅提昇
// 為了優化使用者的體驗時間
// 想改為事件驅動的最終一致性機制, 較為容易更改
// 而不需要改 repo 函數簽名
type TxFactory interface {
	Tx() (driver.Tx, error)
	ContextWithTx(ctx context.Context) (ctxTx context.Context, tx driver.Tx, err error)
}

func NewTxFactory(db *sqlx.DB) *TxFactoryImp {
	return &TxFactoryImp{db: db}
}

type TxFactoryImp struct {
	db *sqlx.DB
}

func (f *TxFactoryImp) Tx() (driver.Tx, error) {
	return f.db.Beginx()
}

func (f *TxFactoryImp) ContextWithTx(ctx context.Context) (ctxTx context.Context, tx driver.Tx, err error) {
	tx, err = f.db.Beginx()
	if err != nil {
		return ctx, nil, err
	}
	ctxTx = context.WithValue(ctx, f.db, tx) // 表示是 這個 db 的 tx 連線
	return ctxTx, tx, nil
}

// ExecWithTxOrDB is used to support that uow.TxFactory's method ContextWithTx
func ExecWithTxOrDB(ctx context.Context, db *sqlx.DB) (sqlExec sqlx.ExtContext) {
	if ctx == nil {
		panic("ctx context.Context is nil")
	}

	// 如果是不同 db 的 tx 連線
	// 則在自己的聚合內, 完成事務交易即可
	// 不和其他聚合合作
	externalTx, ok := ctx.Value(db).(*sqlx.Tx)
	if ok {
		return externalTx
	}
	return db
}

func ReadWithTxOrDB(ctx context.Context, db *sqlx.DB) (sqlExe sqlx.QueryerContext) {
	if ctx == nil {
		panic("ctx context.Context is nil")
	}
	externalTx, ok := ctx.Value(db).(*sqlx.Tx)
	if ok {
		return externalTx
	}
	return db
}

// HandleErrorByRollback
// 參數 err 代表外部錯誤, 預期不等於 nil
// 回傳 nil 代表 rollback 成功
// 依據業務 有不同的錯誤處理方式
// rollback 成功不代表 已經處理 外部錯誤
func HandleErrorByRollback(err error, tx driver.Tx) (rollbackErr error) {
	if err != nil {
		if rollbackErr = tx.Rollback(); rollbackErr != nil {
			return rollbackErr
		}
	}
	return nil
}
