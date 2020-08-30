package helpertype

import (
	"errors"
	"reflect"

	"github.com/morikuni/failure"
	"github.com/rs/zerolog/log"
)

type FieldName = string

func StructFilterZeroValueField(raw interface{}, tagKey string) map[FieldName]interface{} {
	if !isStructType(raw) {
		log.Fatal().Msg("is not struct type")
	}

	v := reflect.ValueOf(raw)
	if v.IsZero() {
		return nil
	}

	var values map[string]interface{}
	var err error

	switch v.Kind() {
	case reflect.Struct:
		values, err = filter(raw, tagKey)
	case reflect.Ptr:
		values, err = filter(v.Elem().Interface(), tagKey)
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
		return isStructType(v.Elem().Interface())
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
		if fieldValue.IsZero() { // filter condition
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
