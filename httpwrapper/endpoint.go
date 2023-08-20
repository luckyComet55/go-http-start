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
	if path[len(path)-1] != '/' {
		path += "/"
	}
	return Endpoint{
		Path:    path,
		Method:  method,
		Handler: handler,
	}
}
