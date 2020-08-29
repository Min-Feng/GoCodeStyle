package domain

import "github.com/morikuni/failure"

const (
	ErrNotFound failure.StringCode = "Not Found"
	ErrDB       failure.StringCode = "DB Layer Failed"
	// ErrValidate failure.StringCode = "ValidateFailed"
	// ErrNil      failure.StringCode = "NilPointer"
)
