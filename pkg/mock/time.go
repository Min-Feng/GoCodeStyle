package mock

import (
	"time"

	"github.com/rs/zerolog/log"
)

//nolint:gochecknoglobals
var timeSpec = []string{
	"01-02",
	"15:04",
	"01-02T15:04",
	"2006-01-02",
	"15:04:05",
	"2006-01-02T15:04",
	"2006-01-02T15:04:05",
	"2006-01-02T15:04:05Z07",
	time.RFC3339,
}

// 若假冒的時間有時區資訊, 只有0時區 UTC 用 Z
// 其他時區要用 +-, 例如 2055-04-01T22:13:47-01:00
func TimeNow(fakeDate string) func() time.Time {
	return func() (fakeTime time.Time) {
		var err error

		for _, spec := range timeSpec {
			fakeTime, err = time.ParseInLocation(spec, fakeDate, time.Local)
			if err != nil {
				continue
			}
			return
		}

		//noinspection GoNilness
		log.Fatal().Msg("Create fake time.Now() failed: " + err.Error())
		return
	}
}
