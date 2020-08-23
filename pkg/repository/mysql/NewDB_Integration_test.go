// +build integration

package mysql_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"ddd/pkg/helper/helpertest/mock"
	"ddd/pkg/repository/mysql"
)

func TestNewDB(t *testing.T) {
	db := mysql.NewDB(&mock.Config().MySQL)
	assert.NotNil(t, db)
}
