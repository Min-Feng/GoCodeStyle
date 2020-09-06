package helpertype

import (
	"errors"
	"reflect"

	"github.com/davecgh/go-spew/spew"
	"github.com/morikuni/failure"
	"github.com/rs/zerolog/log"
)

type FieldName = string

type StructTool struct{}

func (StructTool) FilterZeroValueField(raw interface{}, tagKey string) map[FieldName]interface{} {
	v := reflect.ValueOf(raw)
	if v.IsZero() {
		return map[string]interface{}{}
	}

	if !isStructType(raw) {
		log.Fatal().Msg("is not struct type")
	}

	var values map[string]interface{}
	var err error

	switch v.Kind() {
	case reflect.Struct:
		values, err = filter(raw, tagKey)
	case reflect.Ptr:
		instanceRaw := v.Elem().Interface()
		values, err = filter(instanceRaw, tagKey)
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
		var instanceRaw interface{}

		if v.IsNil() {
			// panic: reflect: call of reflect.Value.Interface on zero Value
			// instanceRaw = reflect.Indirect(v).Interface()
			instanceRaw = ReflectTool{}.NewInstanceValueByPtrValue(v).Interface()
		} else {
			instanceRaw = v.Elem().Interface()
		}

		return isStructType(instanceRaw)
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

		if log.Trace().Enabled() {
			fieldType := structType.Field(i)
			log.Trace().
				Bool("IsZero", fieldValue.IsZero()).
				Msgf("%v %v=%#v", structType.Name(), fieldType.Name, spew.NewFormatter(fieldValue.Interface()))
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
