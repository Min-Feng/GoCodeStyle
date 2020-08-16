package testtool

import (
	"strings"
)

func FormatToRawSQL(prettySQL string) (rawSQL string) {
	for _, s := range []string{"\t", "\n"} {
		prettySQL = strings.ReplaceAll(prettySQL, s, "")
	}
	return prettySQL
}
