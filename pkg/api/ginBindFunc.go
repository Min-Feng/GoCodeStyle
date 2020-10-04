package api

import (
	"reflect"
	"time"

	"github.com/go-playground/validator/v10"
)

func AutoSetTime(fl validator.FieldLevel) bool {
	t, ok := fl.Field().Interface().(time.Time)
	if ok {
		t = time.Now()
		fl.Field().Set(reflect.ValueOf(t))
		return true
	}
	return false
}
