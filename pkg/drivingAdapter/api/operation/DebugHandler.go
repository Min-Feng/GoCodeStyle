package operation

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/morikuni/failure"
	"github.com/rs/zerolog/log"

	"ddd/pkg/domain/basic"
	"ddd/pkg/drivingAdapter/api/shared"
	"ddd/pkg/technical/logger"
)

type DebugHandler struct{}

func (h *DebugHandler) UpdateLogLevel(c *gin.Context) {
	type Payload struct {
		LogLevel   string    `json:"log_level" binding:"required"`
		UpdateTime time.Time `json:"UpdateTime" binding:"AutoSetTime"`
	}

	payload := new(Payload)
	if err := c.ShouldBindJSON(payload); err != nil {
		Err := failure.Translate(err, basic.ErrValidate)
		c.Set("Error", Err)
		log.Error().Msgf("Show ShouldBindJSON error stack%v", Err)
		return
	}
	// log.Info().Msgf("payload=%#v", spew.NewFormatter(payload))

	err := logger.SetGlobal(payload.LogLevel, logger.WriterKindHuman)
	if err != nil {
		Err := failure.Wrap(err)
		c.Set("Error", Err)
		log.Error().Msgf("%v", Err)
		if log.Debug().Enabled() {
			log.Error().Msgf("Show SetGlobal error stack\n%+v", Err)
		}
		return
	}
	log.Info().Str("logLevel", payload.LogLevel).Msg("Set loglevel successfully")

	// 已經執行 return response, 後續的 middleware 就不要再次對 response 進行操作
	c.JSON(http.StatusOK, shared.NormalResponse)
}
