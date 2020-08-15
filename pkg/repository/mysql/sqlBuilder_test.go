package mysql

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

func TestGenericSQLBuilder(t *testing.T) {
	suite.Run(t, new(GenericSQLBuilderTestSuite))
}

type GenericSQLBuilderTestSuite struct {
	suite.Suite
	b GenericSQLBuilder
}

func (ts *GenericSQLBuilderTestSuite) TestIsTheRowExist() {
	actualSQLString, _ := ts.b.IsTheRowExist("member_id", 2, MemberTableName)
	expectSQLString := "SELECT member_id FROM member WHERE member_id = ? FOR UPDATE"
	ts.Assert().Equal(expectSQLString, actualSQLString)
}
