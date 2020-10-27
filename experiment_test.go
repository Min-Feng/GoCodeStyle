// +build experiment

package experiment

import (
	"encoding/json"
	"testing"

	"github.com/davecgh/go-spew/spew"

	"ddd/pkg/technical/logger"
)

func init() {
	logger.DeveloperMode()
}

// 實驗區 測試想法
func TestExperiment(t *testing.T) {
	b := []byte(`{
  "age": 12.2,
  "money": null
}`)

	m := make(map[string]interface{})
	json.Unmarshal(b, &m)

	spew.Dump(m)
}
