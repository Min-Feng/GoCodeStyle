package api

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/gin-gonic/gin"
	"github.com/morikuni/failure"
	"github.com/rs/zerolog/log"

	"ddd/pkg/domain"
	"ddd/pkg/helper/helperlog"
)

type DebugHandler struct{}

func (h *DebugHandler) UpdateLogLevel(c *gin.Context) {
	type Payload struct {
		LogLevel string `json:"log_level"`
	}

	payload := new(Payload)
	if err := c.ShouldBindJSON(payload); err != nil {
		c.Set("Error", failure.Translate(err, domain.ErrValidate))
		c.Next()
		return
	}

	log.Info().Msgf("payload=%#v", spew.NewFormatter(payload))
	helperlog.SetGlobal(payload.LogLevel, helperlog.WriterKindHuman)

	log.Trace().Msg("trace")
	log.Debug().Msg("debug")
	log.Info().Msg("info")
	log.Error().Msg("error")

	// 已經執行 return response, 就不要在後續的 middleware 再次對 response 進行操作
	// c.JSON(http.StatusOK, normalResponse)
}
