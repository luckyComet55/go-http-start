package httpwrapper

import (
	"net/http"
)

type Endpoint struct {
	path       string
	method     httpMethod
	pathParams []string
	handler    http.HandlerFunc
}

func NewEndpoint(path string, method httpMethod, params []string, handler http.HandlerFunc) Endpoint {
	return Endpoint{
		path:       path,
		method:     method,
		pathParams: params,
		handler:    handler,
	}
}
