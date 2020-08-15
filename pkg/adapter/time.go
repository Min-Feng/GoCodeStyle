package adapter

import "time"

// TimeNowFunc 其函數簽名 與 標準庫的 time.Now() 一致
type TimeNowFunc func() time.Time
