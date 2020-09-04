package api

import (
	"net/http"

	"github.com/morikuni/failure"

	"ddd/pkg/domain"
)

const ErrUndefined failure.StringCode = "Undefined Error"

type ResponseCode struct {
	HTTPCode       int
	CustomizedCode int
}

var ErrCodeLookup = map[string]ResponseCode{
	domain.ErrDB.ErrorCode(): {
		HTTPCode:       http.StatusInternalServerError,
		CustomizedCode: 5000,
	},
	domain.ErrNotFound.ErrorCode(): {
		HTTPCode:       http.StatusNotFound,
		CustomizedCode: 4040,
	},
	domain.ErrValidate.ErrorCode(): {
		HTTPCode:       http.StatusBadRequest,
		CustomizedCode: 4000,
	},
	domain.ErrServer.ErrorCode(): {
		HTTPCode:       http.StatusInternalServerError,
		CustomizedCode: 5000,
	},
	ErrUndefined.ErrorCode(): {
		HTTPCode:       http.StatusInternalServerError,
		CustomizedCode: 9999,
	},
}
