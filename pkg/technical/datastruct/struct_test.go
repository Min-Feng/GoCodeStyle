package datastruct_test

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/assert"

	"ddd/pkg/technical/datastruct"
	"ddd/pkg/technical/mock"
)

func TestStructTool_FilterZeroValueField(t *testing.T) {
	// logger.FixBugMode()

	type QueryCondition struct {
		CreatedTime   datastruct.Time `db:"created_time"`
		UserName      string          `db:"user_name"`
		Orders        []string        `db:"order"`
		Age           *int            `db:"age"`
		NullableValue interface{}     `db:"money"`
	}

	tests := []struct {
		name        string
		rawStruct   *QueryCondition
		expectedMap map[datastruct.FieldName]interface{}
	}{
		{
			rawStruct: &QueryCondition{
				UserName: "caesar",
				Orders:   []string{"book", "tea"},
			},
			expectedMap: map[datastruct.FieldName]interface{}{
				"user_name": "caesar",
				"order":     []string{"book", "tea"},
			},
		},
		{
			rawStruct: &QueryCondition{
				CreatedTime: mock.Time("2020-08-23"),
				Orders:      []string{}, // reflect.ValueOf([]string{}).IsZero=false
				UserName:    "caesar",
			},
			expectedMap: map[datastruct.FieldName]interface{}{
				"created_time": mock.Time("2020-08-23"),
				"user_name":    "caesar",
			},
		},
		{
			name:        "All Zero value fields # Struct Not Nil",
			rawStruct:   &QueryCondition{},
			expectedMap: map[string]interface{}{},
		},
		{
			name:        "All Zero value fields # Struct Is Nil",
			rawStruct:   nil,
			expectedMap: map[string]interface{}{},
		},
		{
			name: "field have value but is js null",
			rawStruct: &QueryCondition{
				Age:           (*int)(nil),
				NullableValue: (*float64)(nil), // reflect.ValueOf((*float64)(nil)).IsZero=false
			},
			expectedMap: map[string]interface{}{
				"money": (*float64)(nil),
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			log.Debug().Msgf("\n%v", spew.Sdump(tt.rawStruct))

			actualMap := datastruct.StructTool{}.FilterZeroValueField(tt.rawStruct, "db")
			assert.Equal(t, tt.expectedMap, actualMap)
		})
	}
}
