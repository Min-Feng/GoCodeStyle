package api

import (
	"net/http"

	"ddd/pkg/domain"
)

// Response

var normalResponse = necessaryResponse{
	CustomizedCode: 0,
	Message:        "ok",
}

type necessaryResponse struct {
	CustomizedCode int    `json:"code"`
	Message        string `json:"message"`
}

// Code

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
