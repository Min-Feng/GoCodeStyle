package types

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

type TimeTool struct{}

// timeLayout 格式 只有0時區用 Z
// 其他時區要用 +-, 例如 2055-04-01 22:13:47-01:00
func (TimeTool) Parse(timeLayout string) (t time.Time, err error) {
	for _, spec := range timeSpec {
		t, err = time.ParseInLocation(spec, timeLayout, time.Local)
		if err == nil {
			return t, nil
		}
	}
	return t, err
}
