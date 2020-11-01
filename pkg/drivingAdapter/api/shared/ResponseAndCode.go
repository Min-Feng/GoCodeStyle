package shared

import (
	"net/http"

	"ddd/pkg/domain/basic"
)

// Response

var NormalResponse = NecessaryResponse{
	CustomizedCode: 0,
	Message:        "ok",
}

type NecessaryResponse struct {
	CustomizedCode int    `json:"code"`
	Message        string `json:"message"`
}

// Code

type ResponseCode struct {
	HTTPCode       int
	CustomizedCode int
}

var ErrCodeLookup = map[string]ResponseCode{
	basic.ErrDB.ErrorCode(): {
		HTTPCode:       http.StatusInternalServerError,
		CustomizedCode: 5000,
	},
	basic.ErrNotFound.ErrorCode(): {
		HTTPCode:       http.StatusNotFound,
		CustomizedCode: 4040,
	},
	basic.ErrValidate.ErrorCode(): {
		HTTPCode:       http.StatusBadRequest,
		CustomizedCode: 4000,
	},
	basic.ErrServer.ErrorCode(): {
		HTTPCode:       http.StatusInternalServerError,
		CustomizedCode: 5000,
	},
	ErrUndefined.ErrorCode(): {
		HTTPCode:       http.StatusInternalServerError,
		CustomizedCode: 9999,
	},
}
