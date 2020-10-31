// +build integration

package part

import (
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"

	"ddd/pkg/technical/logger"
	"ddd/pkg/technical/mock"
)

func TestNewRecoverableDB_SyncRequest(t *testing.T) {
	logger.SetGlobal(logger.Disabled, logger.WriterKindHuman)
	// logger.DeveloperMode()

	// 不包含初始化的那次
	// 後續進行 retry init, 會產生 x= 7次 nil DB
	// 所以需要 x+1= 8次 failover action 才可拿到正常的 DB
	maxNilCount := 7

	var count int // 代表 main function 初始化失敗, 回傳 nil
	db := NewRecoverableDB(func() *sqlx.DB {
		if count <= maxNilCount {
			count++
			return nil
		}
		return NewMySQL(&mock.Config.MySQL)
	})

	for i := 1; i <= 50; i++ {
		// 每次呼叫消耗 2次 failover action
		// 一次 avoidExecuteNilDB
		// 一次 avoidCrash
		err := db.forUnitTest()

		if i*2 <= maxNilCount || i*2 <= maxNilCount+1 {
			// 進行 avoidCrash 即使可以復原 DB 狀態
			// 但 panic 已經產生, 依然會回傳錯誤
			// 下次呼叫才會是正常的 DB
			assert.Error(t, err, fmt.Sprintf("i=%v", i))
			continue
		}
		assert.NoError(t, err, fmt.Sprintf("i=%v", i))
	}
}

func TestNewRecoverableDB_AsyncRequest(t *testing.T) {
	logger.SetGlobal(logger.Disabled, logger.WriterKindHuman)
	// logger.DeveloperMode()

	maxNilCount := 7
	var count int
	db := NewRecoverableDB(func() *sqlx.DB {
		if count <= maxNilCount {
			count++
			return nil
		}
		return NewMySQL(&mock.Config.MySQL)
	})

	// 模擬三波瞬間大量請求
	// 每波請求 200 次
	wg := sync.WaitGroup{}
	for j := 0; j < 3; j++ {
		time.Sleep(300 * time.Millisecond)
		for i := 1; i <= 200; i++ {
			wg.Add(1)
			go func() {
				db.forUnitTest()
				wg.Done()
			}()
		}
	}
	wg.Wait()

	err := db.forUnitTest()
	assert.NoError(t, err)
}
