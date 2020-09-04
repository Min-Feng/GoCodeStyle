package api

import (
	"github.com/gin-gonic/gin"
	"github.com/morikuni/failure"
	"github.com/rs/zerolog/log"
)

func ErrorResponseMiddleware(c *gin.Context) {
	c.Next()

	err, ok := c.Get("Error")
	if !ok {
	}

	Err := err.(error)
	log.Error().Err(Err).Send()
	if log.Debug().Enabled() {
		log.Error().Msgf("\n%+v", Err)
	}
	checkoutResponse(c, Err)

	c.Next()
}

var normalResponse = necessaryResponse{
	CustomizedCode: 0,
	Message:        "ok",
}

type necessaryResponse struct {
	CustomizedCode int    `json:"code"`
	Message        string `json:"message"`
}

func checkoutResponse(c *gin.Context, err error) {
	if err == nil {
		return
	}

	causeErr, ok := failure.CodeOf(err)
	if !ok {
		causeErr = ErrUndefined
	}

	ErrString := causeErr.ErrorCode()
	ErrResponse := necessaryResponse{
		CustomizedCode: ErrCodeLookup[ErrString].CustomizedCode,
		Message:        err.Error(),
	}

	c.JSON(ErrCodeLookup[ErrString].HTTPCode, ErrResponse)
}
