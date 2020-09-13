package datastruct_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/suite"

	"ddd/pkg/assistant/datastruct"
	"ddd/pkg/assistant/mock"
)

func TestTime(t *testing.T) {
	suite.Run(t, new(TimeTestSuite))
}

type TimeTestSuite struct {
	suite.Suite
}

func (ts *TimeTestSuite) TestUnmarshalJSON() {
	type Payload struct {
		Time datastruct.Time `json:"datetime"`
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
			err := json.Unmarshal(tt.jsonPayload, payload)

			if tt.name == "failed case" {
				ts.Assert().Error(err, "because there are no 30 days in February")
				return
			}
			ts.Assert().NoError(err)
		})
	}
}

func (ts *TimeTestSuite) TestMarshalJSON() {
	type Payload struct {
		Time datastruct.Time `json:"datetime"`
	}

	tests := []struct {
		name         string
		payload      Payload
		expectedJSON string
	}{
		{
			name:         "Normal Time",
			payload:      Payload{mock.Time("2020-08-16 23:22:55")},
			expectedJSON: `{"datetime":"2020-08-16 23:22:55"}`,
		},
		{
			name:         "Zero Time",
			payload:      Payload{},
			expectedJSON: `{"datetime":""}`,
		},
	}

	for _, tt := range tests {
		tt := tt
		ts.Run(tt.name, func() {
			b, err := json.Marshal(tt.payload)
			ts.Require().NoError(err)

			actualJSON := string(b)
			ts.Assert().Equal(tt.expectedJSON, actualJSON)
		})
	}
}
