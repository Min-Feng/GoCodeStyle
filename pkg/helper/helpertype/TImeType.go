package helpertype

import (
	"bytes"
	"database/sql"
	"database/sql/driver"
	"time"

	"github.com/morikuni/failure"
)

//nolint:gochecknoglobals
var (
	TimeLayoutJSON        = "2006-01-02 15:04:05"
	TimeLayoutSQLDatetime = "2006-01-02 15:04:05"
	// TimeLayoutSQLDate     = "2006-01-02"
)

// helpertype.Time 定義新的時間型別, 是為了處理以下兩種情況
//
// 1. 資料庫欄位 allow NULL datetime 型別
//
// 2. 與前端溝通的 time 格式, 如何進行 Marshal, Unmarshal
type Time struct {
	time.Time
}

/*
 當 helpertype.Time 為零值, 則寫入資料庫時, 填寫 NULL
 若 資料庫 datetime 欄位為 NOT NULL
 則此方法必須改寫, 或創造另一個 helpertype.Time

 在同一個專案或同一家公司, 不要存在兩種 sql.datetime 限制

 這裡使用 value receiver 來實現 interface: sql.Valuer

 當序列化資料時, 轉換儲存媒介
 資料型別使用 pointer 對另一個儲存媒介是沒意義的
 我希望使用 value 來表示資料型別
 所以使用了 value receiver 來實現 interface: sql.Valuer

 若使用 pointer receiver 來實現 interface
 那麼資料型別就必須使用 pointer 才可以滿足 interface 的規則
 因為 value instance 無法得知 pointer receiver 所擁有的方法集合
 https://golang.org/ref/spec#Method_sets
*/
func (t Time) Value() (driver.Value, error) {
	if t.IsZero() {
		return nil, nil
	}
	return t.sqlDatetime(), nil
}

// squirrel 使用 sq.Eq{} 進行 ToSql() 的時候
// 會使用型別所實現的 interface: driver.Value
//
// 但若使用 squirrel Insert Value 進行 ToSql() 的時候
// 則是使用型別所實現的 interface: fmt.Stringer
//
// 後來發現, 似乎只有 sq.DebugSqlizer 才會轉為 fmt.Stringer
// https://github.com/Masterminds/squirrel/issues/260
//
// 最後都轉化為 ToSql() 的 args
func (t Time) String() string {
	return t.sqlDatetime()
}

// sqlDatetime 主要用在 sql where 的地方, 且該欄位在資料庫為 datetime 型別
// 因為不知道底層是怎麼轉換 time.Time
// 所以採用顯式轉換字串格式, 避免輸出的格式不如預期, 造成 sql index 失效
func (t Time) sqlDatetime() string {
	return t.Format(TimeLayoutSQLDatetime)
}

func (t Time) sqlUnix() int64 {
	return t.Unix()
}

func (t *Time) Scan(src interface{}) error {
	if src == nil {
		return nil
	}

	var helperTime sql.NullTime

	err := helperTime.Scan(src)
	if err != nil {
		return failure.Wrap(err, failure.Message("helpertype.Time sql Scan"))
	}

	t.Time = helperTime.Time
	return nil
}

// 標準庫 time.Time 其 json.Marshal 的時間格式: "2020-08-16T23:22:55+08:00"
// 改成專案規範的格式: "2020-08-16 23:22:55"
// 且若 Time 為零值, 則輸出到 json 的格式為空字串
func (t Time) MarshalJSON() ([]byte, error) {
	//nolint:gomnd
	b := make([]byte, 0, len(TimeLayoutJSON)+2)
	b = append(b, '"')
	if !t.IsZero() {
		b = t.AppendFormat(b, TimeLayoutJSON)
	}
	b = append(b, '"')
	return b, nil
}

// 標準庫 time.Time 其 json.Unmarshal 可接受的時間格式, 只有 time.RFC3339
// 為了更加方便使用, 重新定義多種 time spec
// 且收到 空字串"" 或 "null", 則會視為 helpertype.Time 零值
//
// 在標準庫 go1.13, 理論上只接受 "null" 為 time.Time 零值
// 但標準庫有 bug, 無法接受 "null", 因為沒有考慮到雙引號
func (t *Time) UnmarshalJSON(data []byte) (err error) {
	timeString := string(bytes.Trim(data, `"`))

	if timeString == "null" || timeString == "" || timeString == "nil" {
		return nil
	}

	t.Time, err = TimeTool{}.Parse(timeString)
	if err != nil {
		return failure.Wrap(err, failure.Message("helpertype.Time UnmarshalJSON"))
	}
	return nil
}
