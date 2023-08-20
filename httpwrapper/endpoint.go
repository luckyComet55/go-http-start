package httpwrapper

import (
	"net/http"
)

type Endpoint struct {
	Path    string
	Method  httpMethod
	Handler http.HandlerFunc
}

func NewEndpoint(path string, method httpMethod, handler http.HandlerFunc) Endpoint {
	return Endpoint{
		Path:    path,
		Method:  method,
		Handler: handler,
	}
}
