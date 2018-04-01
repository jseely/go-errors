package errors

import (
	"fmt"
)

func New(text string) error {
	return &stringError{text}
}

type stringError struct {
	s string
}

func (e *stringError) Error() string {
	return s.s
}
