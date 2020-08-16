package adapter_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/suite"

	"ddd/pkg/adapter"
	"ddd/pkg/loghelper"
	"ddd/pkg/mock"
)

func init() {
	loghelper.UnitTestSetting()
}

func TestTime(t *testing.T) {
	suite.Run(t, new(TimeTestSuite))
}

type TimeTestSuite struct {
	suite.Suite
}

func (ts *TimeTestSuite) TestUnmarshalJSON() {
	type Payload struct {
		Time adapter.Time `json:"datetime"`
	}

	tests := []struct {
		name        string
		jsonPayload []byte
	}{
		{jsonPayload: []byte(`{"age":18}`)},
		{jsonPayload: []byte(`{"datetime":"2020-08-16 10:27"}`)},
		{jsonPayload: []byte(`{"datetime":"2020-08-16T23:22"}`)},
		{jsonPayload: []byte(`{"datetime":   "2020-08-16T23:22"      }`)},
		{jsonPayload: []byte(`{"datetime":"2020-08-16"}`)},
		{jsonPayload: []byte(`{"datetime":""}`)},
		{jsonPayload: []byte(`{"datetime":"null"}`)},
		{jsonPayload: []byte(`{"datetime":"2020-08-16 23:22:55"}`)},
		{jsonPayload: []byte(`{"datetime":"2020-08-16T23:22:55"}`)},
		{jsonPayload: []byte(`{"datetime":"2020-08-16T23:22:55+08:00"}`)},
		{jsonPayload: []byte(`{"datetime":"nil"}`)},
		{name: "failed case", jsonPayload: []byte(`{"datetime":"2020-02-30 23:22:55"}`)},
	}

	for _, tt := range tests {
		tt := tt
		ts.Run(tt.name, func() {
			payload := new(Payload)
			payload.Time = adapter.Time{}
			err := json.Unmarshal(tt.jsonPayload, payload)

			if tt.name == "failed case" {
				ts.Assert().Error(err)
				return
			}
			ts.Assert().NoError(err)
		})
	}
}

func (ts *TimeTestSuite) TestMarshalJSON() {
	type Payload struct {
		Time adapter.Time `json:"datetime"`
	}
	payload := Payload{}
	payload.Time = adapter.Time{mock.NewTimeNowFunc("2020-08-16 23:22:55")()}

	b, err := json.Marshal(payload)
	ts.Assert().NoError(err)
	actualTimeFormat := string(b)

	expectedTimeFormat := `{"datetime":"2020-08-16 23:22:55"}`
	ts.Assert().Equal(expectedTimeFormat, actualTimeFormat)
}
