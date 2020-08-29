package mysql

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/morikuni/failure"
	"github.com/rs/zerolog/log"
	"golang.org/x/sync/singleflight"

	"ddd/pkg/domain"
)

func NewRecoverableDB(newDB func() *sqlx.DB) *RecoverableDB {
	return &RecoverableDB{
		DB:    newDB(),
		newDB: newDB,
	}
}

// 正常流程, 完全不會使用到 failover 的機制
// 若遭遇故障, 期望達成 3 個目標
//
// 1. 若 failover 處理失敗, 還會進行下一次
//
// 2. 循序請求的情境
// 每次都進行故障處理
//
// 3. 瞬間大量請求的情境
// 只執行必要次數的 failover
//
// ex:
//
// 假設進行 8 次 failover action 資料庫就可以復原
// 即使瞬間有 1000 次查詢
// failover 只會執行 8 次的
// 不會進行多餘的次數
type RecoverableDB struct {
	*sqlx.DB
	newDB         func() *sqlx.DB
	failover      singleflight.Group
	failoverTimes int
}

func (reDB *RecoverableDB) forUnitTest() (Err error) {
	defer func() {
		// 只有在當前函數發生 panic
		// 呼叫 recover 才可以捕捉到 panic message
		// 若是重構為成函數來呼叫
		// 要注意該函數不可包含 if panicMsg := recover(); panicMsg != nil {
		if panicMsg := recover(); panicMsg != nil {
			err := fmt.Errorf("%v", panicMsg)
			Err = failure.Translate(err, domain.ErrDB)
			reDB.avoidCrash()
		}
	}()

	if reDB.DB == nil {
		reDB.avoidExecuteNilDB()
	}

	return reDB.DB.Ping()
}

// func (reDB *RecoverableDB) Get(dest interface{}, query string, args ...interface{}) (err error) {
// 	panic("")
// }

func (reDB *RecoverableDB) avoidCrash() {
	reDB.failoverAction("panic occurs", func() {
		reDB.close()
		reDB.retryInit()
	})
}

// avoidExecuteNilDB is for if main function execute init DB failed and only if return DB is nil
func (reDB *RecoverableDB) avoidExecuteNilDB() {
	reDB.failoverAction("detect DB is nil", reDB.retryInit)
}

func (reDB *RecoverableDB) failoverAction(failoverReason string, action func()) {
	// 瞬間大量請求的情境, share 為 true
	// 循序請求的情境, share 為 false
	_, _, shared := reDB.failover.Do(failoverReason, func() (interface{}, error) {
		log.Info().Str("Failover", failoverReason).Send()
		reDB.failoverTimes++
		action()
		return nil, nil
	})
	if log.Debug().Enabled() {
		log.Debug().Bool("share", shared).Send()
	}
}

func (reDB *RecoverableDB) retryInit() {
	log.Info().Str("Failover", fmt.Sprintf("%vth action start: retry to init db", reDB.failoverTimes)).Send()

	reDB.DB = reDB.newDB()
	if reDB.DB == nil {
		log.Error().Str("Failover", fmt.Sprintf("%vth action end: init mysql db failed", reDB.failoverTimes)).Send()
		return
	}
	log.Info().Str("Failover", fmt.Sprintf("%vth action end: init mysql db successfully", reDB.failoverTimes)).Send()
}

func (reDB *RecoverableDB) close() {
	if reDB.DB == nil {
		return
	}
	log.Info().Str("Failover", fmt.Sprintf("%vth action start: close db", reDB.failoverTimes)).Send()

	if err := reDB.DB.Close(); err != nil {
		log.Error().Err(err).Str("Failover", fmt.Sprintf("%vth action end: close db", reDB.failoverTimes)).Send()
		return
	}
	log.Info().Str("Failover", fmt.Sprintf("%vth action end: close db successfully", reDB.failoverTimes)).Send()
}
