// +build integration

package mysql_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"ddd/pkg/repository/mysql"
	"ddd/pkg/testtool"
)

func TestNewDB(t *testing.T) {
	cfg := testtool.Config()
	db := mysql.NewDB(&cfg.Mysql)
	assert.NotNil(t, db)
}
