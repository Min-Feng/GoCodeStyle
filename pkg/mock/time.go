package mock

import (
	"time"

	"github.com/rs/zerolog/log"

	"ddd/pkg/adapter"
)

func NewTimeNowFunc(fakeTime string) adapter.TimeNowFunc {
	return func() time.Time {
		t, err := adapter.TimeParse(fakeTime)
		if err != nil {
			log.Fatal().Msgf("New fakeTimeNow function failed: %v\n%+[1]v", err)
		}
		return t
	}
}
