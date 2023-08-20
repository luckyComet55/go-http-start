package httpwrapper

import (
	"net/http"
)

type Endpoint struct {
	Path       string
	Method     httpMethod
	PathParams []string
	Handler    http.HandlerFunc
}

func NewEndpoint(path string, method httpMethod, params []string, handler http.HandlerFunc) Endpoint {
	return Endpoint{
		Path:       path,
		Method:     method,
		PathParams: params,
		Handler:    handler,
	}
}
