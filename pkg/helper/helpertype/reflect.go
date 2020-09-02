package helpertype

import (
	"reflect"

	"github.com/rs/zerolog/log"
)

type ReflectTool struct{}

// NewInstanceValueByPtrValue can help to be able to continue any workflow, when v.IsNil() is true
//
// example: raw is nil
//
// var raw *int
// v := reflect.ValueOf(raw)
// v.Kind() is reflect.Ptr
//
// DEPRECATED: because the function purpose just as reflect.Indirect
func (ReflectTool) NewInstanceValueByPtrValue(v reflect.Value) reflect.Value {
	if v.Kind() != reflect.Ptr {
		log.Fatal().Msg("ReflectTool.NewInstanceValueByPtrValue: parameter kind is not reflect.Ptr")
		return reflect.Value{}
	}

	// just as:
	// var data int
	// v := reflect.ValueOf(&data)

	// OriginType:=reflect.TypeOf(data)
	OriginType := v.Type().Elem()

	// OriginPointerValue:=reflect.ValueOf(&data)
	OriginPointerValue := reflect.New(OriginType)

	// OriginInstanceValue:=reflect.ValueOf(data)
	OriginInstanceValue := reflect.Indirect(OriginPointerValue)
	return OriginInstanceValue
}
