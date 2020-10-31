package types

import (
	"reflect"

	"github.com/rs/zerolog/log"
)

type ReflectTool struct{}

// NewInstanceValueByPtrValue when v.IsNil() is true, can help to be able to continue any workflow.
// Example: avoid scenario that panic: reflect: call of reflect.Value.Interface on zero Value
//
// v.Kind() is reflect.Ptr and v.IsNil() is true
func (ReflectTool) NewInstanceValueByPtrValue(v reflect.Value) reflect.Value {
	if v.Kind() != reflect.Ptr {
		log.Fatal().Msg("ReflectTool.NewInstanceValueByPtrValue: parameter kind is not reflect.Ptr")
		return reflect.Value{}
	}

	// if v.IsNil() is false, just like as:
	// var data int
	// OriginType:=reflect.TypeOf(data)
	// OriginPointerValue:=reflect.ValueOf(&data)
	// OriginInstanceValue:=reflect.ValueOf(data)

	OriginType := v.Type().Elem()                               // find type by pointer
	OriginPointerValue := reflect.New(OriginType)               // new pointer by type
	OriginInstanceValue := reflect.Indirect(OriginPointerValue) // new instance by pointer

	return OriginInstanceValue
}
