package mysql_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"ddd/pkg/repository/mysql"
	"ddd/pkg/testtool"
)

func TestGenericSQLBuilder(t *testing.T) {
	suite.Run(t, new(GenericSQLBuilderTestSuite))
}

type GenericSQLBuilderTestSuite struct {
	suite.Suite
	gSQL mysql.GenericSQLBuilder
}

func (ts *GenericSQLBuilderTestSuite) TestIsTheRowExist() {
	actualSQLString, _ := ts.gSQL.IsTheRowExist("member_id", 2, mysql.MemberTableName)
	expectedSQLString := testtool.FormatToRawSQL(`
			SELECT member_id 
			FROM member 
			WHERE member_id = ? 
			FOR UPDATE`)
	ts.Assert().Equal(expectedSQLString, actualSQLString)
}
