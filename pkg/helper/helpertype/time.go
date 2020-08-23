package helpertype

import "time"

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

// TimeNowFunc
// 給標準庫 time.Now() 一個型別名稱
// 為了可以更清楚知道, 此變數需要的是 time.Now() function
type TimeNowFunc = func() time.Time

// timeLayout 格式 只有0時區用 Z
// 其他時區要用 +-, 例如 2055-04-01 22:13:47-01:00
func TimeParse(timeLayout string) (t time.Time, err error) {
	for _, spec := range timeSpec {
		t, err = time.ParseInLocation(spec, timeLayout, time.Local)
		if err == nil {
			return t, nil
		}
	}
	return t, err
}
