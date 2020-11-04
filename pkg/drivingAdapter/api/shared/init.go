package api

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func Init() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		// v.RegisterCustomTypeFunc()
		v.RegisterValidation("AutoSetTime", AutoSetTime)
	}
}

func init() {
	Init()
}
