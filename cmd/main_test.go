package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/morikuni/failure"
	"github.com/rs/zerolog/log"

	"ddd/pkg/adapter"
)

func TestMain(m *testing.M) {
	var err error
	var data = []byte(`{"created_time":"2200-04-01 23:05:01"}`)

	payload := struct {
		T adapter.Time `json:"created_time"`
	}{}
	err = json.Unmarshal(data, &payload)
	if err != nil {
		log.Error().Msgf("%structValue\n%+[1]structValue", err)
	}
	spew.Dump(payload)

	validPayload, _ := GetValidFieldFromStruct(payload, "json")
	spew.Dump(validPayload)
}

type FieldName = string

func GetValidFieldFromStruct(data interface{}, tagKey string) (map[FieldName]interface{}, error) {
	v := reflect.ValueOf(data)

	switch v.Kind() {
	case reflect.Struct:
		values, err := filterStruct(v, tagKey)
		if err != nil {
			return nil, err
		}
		return values, nil
	case reflect.Ptr:
		values, err := filterStruct(v.Elem().Interface(), tagKey)
		if err != nil {
			return nil, err
		}
		return values, nil
	}

	return nil, failure.Wrap(fmt.Errorf("data type is not struct"))
}

func filterStruct(data interface{}, tagKey string) (map[FieldName]interface{}, error) {
	structValue := reflect.ValueOf(data)
	structType := structValue.Type()
	fieldNum := structValue.NumField()
	values := make(map[string]interface{}, fieldNum)

	for i := 0; i < fieldNum; i++ {
		fieldValue := structValue.Field(i)
		if fieldValue.IsZero() {
			continue
		}

		fieldType := structType.Field(i)
		fieldTagName, ok := fieldType.Tag.Lookup(tagKey)
		if !ok {
			return nil, fmt.Errorf("not found tag key")
		}
		values[fieldTagName] = fieldValue.Interface()
	}
	return values, nil
}
