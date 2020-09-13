package datastruct

import (
	"bytes"
	stdJSON "encoding/json"
	"regexp"
	"strings"

	"github.com/rs/zerolog/log"
	"github.com/tdewolff/minify/v2"
	"github.com/tdewolff/minify/v2/json"
)

type StringTool struct{}

func (StringTool) ToRawJSON(prettyJSON string) string {
	m := minify.New()
	m.AddFuncRegexp(regexp.MustCompile("[/+]json$"), json.Minify)
	rawJSON, err := m.String("application/json", prettyJSON)
	if err != nil {
		log.Fatal().Msgf("ToRawJSON: %v", err)
	}
	return rawJSON
}

func (StringTool) ToPrettyJSON(rawJSON []byte) string {
	var prettyJSON bytes.Buffer
	if err := stdJSON.Indent(&prettyJSON, rawJSON, "", "  "); err != nil {
		log.Fatal().Msgf("ToPrettyJSON", err)
	}
	return prettyJSON.String()
}

func (StringTool) ToRawSQL(prettySQL string) (rawSQL string) {
	rawSQL = prettySQL
	for _, s := range []string{"\t", "\n"} {
		rawSQL = strings.ReplaceAll(rawSQL, s, "")
	}
	return rawSQL
}
