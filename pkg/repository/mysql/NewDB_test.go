// +build integration

package mysql

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"ddd/pkg/loghelper"
)

func init() {
	loghelper.ForUnitTest()
}

func TestNewDB(t *testing.T) {
	loghelper.ForDeveloper()
	cfg := &Config{
		User:        "root",
		Password:    "1234",
		Host:        "127.0.0.1",
		Port:        "3306",
		Database:    "GoCodeStyle",
		MaxConn:     10,
		MaxIdleConn: 5,
	}

	db := NewDB(cfg)
	assert.NotNil(t, db)
}
