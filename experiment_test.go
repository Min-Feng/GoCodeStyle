package GoCodeStyle

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"

	"ddd/pkg/loghelper"
)

func init() {
	loghelper.DevelopSetting()
}

func TestExperiment(t *testing.T) {
	var err error

	// todo: code

	assert.NoError(t, err)
	spew.Dump("")
}
