package helpertype

import (
	"errors"
	"reflect"

	"github.com/morikuni/failure"
	"github.com/rs/zerolog/log"
)

type FieldName = string

type StructTool struct{}

func (StructTool) FilterZeroValueField(raw interface{}, tagKey string) map[FieldName]interface{} {
	if !isStructType(raw) {
		log.Fatal().Msg("is not struct type")
	}

	v := reflect.ValueOf(raw)
	if v.IsZero() {
		return map[string]interface{}{}
	}

	var values map[string]interface{}
	var err error

	switch v.Kind() {
	case reflect.Struct:
		values, err = filter(raw, tagKey)
	case reflect.Ptr:
		var originValue interface{}
		if v.IsNil() {
			originValue = reflect.Indirect(v)
		} else {
			originValue = v.Elem().Interface()
		}
		values, err = filter(originValue, tagKey)
	}

	if err != nil {
		log.Fatal().Msgf("Struct filter zero value field failed: %v\n%+[1]v", err)
	}
	return values
}

func isStructType(raw interface{}) bool {
	v := reflect.ValueOf(raw)
	switch v.Kind() {
	case reflect.Struct:
		return true
	case reflect.Ptr:
		var originValue interface{}

		if v.IsNil() {
			originValue = reflect.Indirect(v)
		} else {
			originValue = v.Elem().Interface()
		}

		return isStructType(originValue)
	}
	return false
}

func filter(raw interface{}, tagKey string) (map[FieldName]interface{}, error) {
	structValue := reflect.ValueOf(raw)
	structType := structValue.Type()
	fieldNum := structValue.NumField()
	values := make(map[string]interface{}, fieldNum)

	for i := 0; i < fieldNum; i++ {
		fieldValue := structValue.Field(i)

		if log.Debug().Enabled() {
			fieldType := structType.Field(i)
			log.Debug().
				Str("FieldName", fieldType.Name).
				Bool("IsZero", fieldValue.IsZero()).
				Msg(structType.Name())
		}

		if fieldValue.IsZero() { // filter condition
			fieldValue.IsValid()
			continue
		}

		fieldType := structType.Field(i)
		fieldTagName, ok := fieldType.Tag.Lookup(tagKey)
		if !ok {
			return nil, failure.Wrap(errors.New("not found tag key"))
		}
		values[fieldTagName] = fieldValue.Interface()
	}
	return values, nil
}
