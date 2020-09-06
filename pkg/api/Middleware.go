package api

import (
	"github.com/gin-gonic/gin"
	"github.com/morikuni/failure"
)

const ErrUndefined failure.StringCode = "Undefined Error"

func ErrorResponseMiddleware(c *gin.Context) {
	c.Next()

	err, ok := c.Get("Error")
	if !ok {
		return
	}

	Err := err.(error)
	checkoutResponse(c, Err)
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
