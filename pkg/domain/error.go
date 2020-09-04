package domain

import "github.com/morikuni/failure"

const (
	ErrNotFound failure.StringCode = "Not Found"
	ErrDB       failure.StringCode = "DB Layer Failed"
	ErrValidate failure.StringCode = "Validate Failed"
	ErrServer   failure.StringCode = "Server Error"
	// ErrNil      failure.StringCode = "NilPointer"
)
