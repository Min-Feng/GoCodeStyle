package mysql_test

import (
	"testing"

	sq "github.com/Masterminds/squirrel"
	"github.com/stretchr/testify/suite"

	"ddd/pkg/helper/helperlog"
	"ddd/pkg/helper/helpertest"
	"ddd/pkg/helper/helpertest/mock"
	"ddd/pkg/helper/helpertype"
	"ddd/pkg/repository/mysql"
)

func TestGenericSQLBuilder(t *testing.T) {
	suite.Run(t, new(GenericSQLBuilderTestSuite))
}

type GenericSQLBuilderTestSuite struct {
	suite.Suite
	gSQL mysql.GenericSQLBuilder
}

func (ts *GenericSQLBuilderTestSuite) TestIsTheRowExist() {
	b := ts.gSQL.IsTheRowExist("member_id", 2, "myTable")
	actualNamedSQL := sq.DebugSqlizer(b)
	expectedSQL := helpertest.FormatToRawSQL(`
		SELECT member_id 
		FROM myTable 
		WHERE member_id = '2' 
		FOR UPDATE`)
	ts.Assert().Equal(expectedSQL, actualNamedSQL)
}

func (ts *GenericSQLBuilderTestSuite) TestTimeRange() {
	helperlog.DeveloperMode()
	timeFieldName := "created_time"

	tests := []struct {
		name             string
		startTime        interface{}
		endTime          interface{}
		expectedNamedSQL string
		expectedSQL      string
	}{
		{
			name:             "Have End StandardTime",
			startTime:        helpertype.Time{mock.StandardTime("2020-08-19 19:43:00")},
			endTime:          helpertype.Time{mock.StandardTime("2020-08-21 00:00:00")},
			expectedNamedSQL: "(created_time >= '2020-08-19 19:43:00' AND created_time <= '2020-08-21 00:00:00')",
			expectedSQL:      "(created_time >= ? AND created_time <= ?)",
		},
		{
			name:             "No End StandardTime",
			startTime:        helpertype.Time{mock.StandardTime("2020-08-19 19:43:00")},
			endTime:          nil,
			expectedNamedSQL: "created_time >= '2020-08-19 19:43:00'",
			expectedSQL:      "created_time >= ?",
		},
		{
			name:             "No Start StandardTime",
			startTime:        nil,
			endTime:          helpertype.Time{mock.StandardTime("2020-08-19 19:43:00")},
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
