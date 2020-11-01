package shared_test

import (
	"testing"

	sq "github.com/Masterminds/squirrel"
	"github.com/stretchr/testify/suite"

	"ddd/pkg/infra/shared"
	"ddd/pkg/technical/logger"
	"ddd/pkg/technical/mock"
	"ddd/pkg/technical/types"
)

func TestGenericSQLBuilder(t *testing.T) {
	suite.Run(t, new(GenericSQLBuilderTestSuite))
}

type GenericSQLBuilderTestSuite struct {
	suite.Suite
	gSQL shared.GenericSQLBuilder
}

func (ts *GenericSQLBuilderTestSuite) TestIsTheRowExist() {
	b := ts.gSQL.SelectForUpdate("member_id", 2, "myTable")
	actualNamedSQL := sq.DebugSqlizer(b)
	expectedSQL := types.StringTool{}.ToRawSQL(`
SELECT member_id 
FROM myTable 
WHERE member_id = '2' 
FOR UPDATE
`)
	ts.Assert().Equal(expectedSQL, actualNamedSQL)
}

func (ts *GenericSQLBuilderTestSuite) TestTimeRange() {
	logger.DeveloperMode()
	timeFieldName := "created_time"

	tests := []struct {
		name             string
		startTime        interface{}
		endTime          interface{}
		expectedNamedSQL string
		expectedSQL      string
	}{
		{
			name:             "Have End StdTime",
			startTime:        types.Time{mock.StdTime("2020-08-19 19:43:00")},
			endTime:          types.Time{mock.StdTime("2020-08-21 00:00:00")},
			expectedNamedSQL: "(created_time >= '2020-08-19 19:43:00' AND created_time <= '2020-08-21 00:00:00')",
			expectedSQL:      "(created_time >= ? AND created_time <= ?)",
		},
		{
			name:             "No End StdTime",
			startTime:        types.Time{mock.StdTime("2020-08-19 19:43:00")},
			endTime:          nil,
			expectedNamedSQL: "created_time >= '2020-08-19 19:43:00'",
			expectedSQL:      "created_time >= ?",
		},
		{
			name:             "No QuicklyStart StdTime",
			startTime:        nil,
			endTime:          types.Time{mock.StdTime("2020-08-19 19:43:00")},
			expectedNamedSQL: "created_time <= '2020-08-19 19:43:00'",
			expectedSQL:      "created_time <= ?",
		},
	}

	for _, tt := range tests {
		tt := tt
		ts.Run(tt.name, func() {
			b := ts.gSQL.TimeRange(timeFieldName, tt.startTime, tt.endTime)

			actualNamedSQL := sq.DebugSqlizer(b)
			ts.Assert().Equal(tt.expectedNamedSQL, actualNamedSQL)

			actualSQL, _, _ := b.ToSql()
			ts.Assert().Equal(tt.expectedSQL, actualSQL)
		})
	}
}
