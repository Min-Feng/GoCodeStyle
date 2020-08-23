package helpertype_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"ddd/pkg/helper/helpertest/mock"
	"ddd/pkg/helper/helpertype"
)

func TestStructFilterZeroValueField(t *testing.T) {
	type QueryCondition struct {
		CreatedTime helpertype.Time `db:"created_time"`
		UserName    string          `db:"user_name"`
		Orders      []string        `db:"order"`
	}

	tests := []struct {
		name        string
		rawStruct   QueryCondition
		expectedMap map[helpertype.FieldName]interface{}
	}{
		{
			rawStruct: QueryCondition{
				UserName: "caesar",
				Orders:   []string{"book", "tea"},
			},
			expectedMap: map[helpertype.FieldName]interface{}{
				"user_name": "caesar",
				"order":     []string{"book", "tea"},
			},
		},
		{
			rawStruct: QueryCondition{
				CreatedTime: mock.CustomizedTime("2020-08-23"),
				UserName:    "caesar",
			},
			expectedMap: map[helpertype.FieldName]interface{}{
				"created_time": mock.CustomizedTime("2020-08-23"),
				"user_name":    "caesar",
			},
		},
		{
			name:        "All Zero value fields",
			rawStruct:   QueryCondition{},
			expectedMap: nil,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			actualMap := helpertype.StructFilterZeroValueField(tt.rawStruct, "db")
			assert.Equal(t, tt.expectedMap, actualMap)
		})
	}
}
