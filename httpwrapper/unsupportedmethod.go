package httpwrapper

import "fmt"

type unsupportedMethodError struct {
	m string
}

func newUnsupportedMethodError(method string) unsupportedMethodError {
	return unsupportedMethodError{
		m: method,
	}
}

func (err unsupportedMethodError) Error() string {
	return fmt.Sprintf("method %s is not supported yet\n", err.m)
}
