package basic

import "github.com/morikuni/failure"

const (
	ErrNotFound failure.StringCode = "Not Found Data"
	ErrDB       failure.StringCode = "DB Layer Failed"
	ErrValidate failure.StringCode = "Validate Parameter Failed"
	ErrServer   failure.StringCode = "Server Error"
)
