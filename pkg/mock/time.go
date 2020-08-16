package mock

import (
	"time"

	"github.com/rs/zerolog/log"

	"ddd/pkg/adapter"
)

func NewTimeNowFunc(fakeDate string) adapter.TimeNowFunc {
	return func() (fakeTime time.Time) {
		t, err := adapter.TimeParse(fakeDate)
		if err != nil {
			log.Fatal().Msgf("New fakeTimeNow function failed: %v\n%+[1]v", err)
		}
		return t
	}
}
