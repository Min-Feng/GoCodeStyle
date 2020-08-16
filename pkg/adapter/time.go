package adapter

import (
	"bytes"
	"database/sql"
	"database/sql/driver"
	"time"

	"github.com/morikuni/failure"
)

// TimeNowFunc 其函數簽名 與 標準庫的 time.Now() 一致
// 單純為了給 標準庫的 time.Now() 一個名字
// 進行依賴注入時, 寫在參數的位置可以更清楚
type TimeNowFunc func() time.Time

// adapter.Time 是為了處理以下兩種情況
// 1. 資料庫欄位 allow NULL datetime 型別
// 2. 與前端溝通的 time 格式
type Time struct {
	time.Time
}

// 當 adapter.Time 為零值, 則寫入資料庫時, 填寫 NULL
// 若 資料庫 datetime 欄位為 NOT NULL
// 則此方法必須改寫, 或創造另一個 adapter.Time
// 在同一個專案或同一家公司, 不要存在兩種 datetime 型別限制
//
// 這裡使用 value receiver
// 是因為 value instance 無法得知 pointer receiver 所擁有的方法集合
func (t Time) Value() (driver.Value, error) {
	if t.IsZero() {
		return nil, nil
	}
	return t.Time, nil
}

func (t *Time) Scan(src interface{}) error {
	if src == nil {
		return nil
	}

	var helper sql.NullTime

	err := helper.Scan(src)
	if err != nil {
		return failure.Wrap(err, failure.Message("adapter.Time sql Scan"))
	}

	t.Time = helper.Time
	return nil
}

// SQLDatetime 主要用在 sql where 的地方, 且該欄位在資料庫為 datetime 型別
// 因為不知道底層是怎麼轉換 time.Time
// 所以採用顯式轉換字串格式, 避免輸出的格式不如預期, 造成 sql index 失效
func (t *Time) SQLDatetime() string {
	return t.Format("2006-01-02 15:04:05")
}

// 標準庫 time.Time 其 json.Marshal 的時間格式: "2020-08-16T23:22:55+08:00"
// 改成專案規範的格式: "2020-08-16 23:22:55"
func (t Time) MarshalJSON() ([]byte, error) {
	layout := "2006-01-02 15:04:05"
	//nolint:gomnd
	b := make([]byte, 0, len(layout)+2)
	b = append(b, '"')
	b = t.AppendFormat(b, layout)
	b = append(b, '"')
	return b, nil
}

// 標準庫 time.Time 其 json.Unmarshal 可接受的時間格式, 只有 time.RFC3339
// 為了更加方便使用, 重新定義多種 time spec
// 且收到 空字串"" 或 "null", 則會視為 adapter.Time 零值
//
// 在標準庫 go1.13, 理論上只接受 "null" 為 time.Time 零值
// 但標準庫有 bug, 無法接受 "null", 因為沒有考慮到雙引號
func (t *Time) UnmarshalJSON(data []byte) (err error) {
	timeString := string(bytes.Trim(data, `"`))

	if timeString == "null" || timeString == "" || timeString == "nil" {
		return nil
	}

	t.Time, err = TimeParse(timeString)
	if err != nil {
		return failure.Wrap(err, failure.Message("adapter.Time UnmarshalJSON"))
	}
	return nil
}

// timeString 格式 只有0時區用 Z
// 其他時區要用 +-, 例如 2055-04-01T22:13:47-01:00
func TimeParse(timeString string) (t time.Time, err error) {
	for _, spec := range timeSpec {
		t, err = time.ParseInLocation(spec, timeString, time.Local)
		if err == nil {
			return t, nil
		}
	}
	return t, err
}

//nolint:gochecknoglobals
var timeSpec = []string{
	"2006-01-02 15:04:05",
	"2006-01-02",
	"2006-01-02 15:04",
	"2006-01-02 15:04:05Z07:00",
	"2006-01-02T15:04",
	"2006-01-02T15:04:05",
	time.RFC3339,
	// "15:04:05",
	// "15:04",
}
