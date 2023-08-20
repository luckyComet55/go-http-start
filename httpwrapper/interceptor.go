package httpwrapper

import (
	"fmt"
	"net/http"
)

type interceptor struct {
	path     string
	handlers map[httpMethod]Endpoint
}

func (i interceptor) intercept(w http.ResponseWriter, r *http.Request) {
	method := httpMethod(r.Method)
	fmt.Println(method)
	endpoint, ok := i.handlers[method]
	var runtimeHandler http.HandlerFunc
	if !ok {
		runtimeHandler = http.NotFound
	} else {
		runtimeHandler = endpoint.Handler
	}
	runtimeHandler(w, r)
}

func newInterceptor(path string) interceptor {
	return interceptor{
		path:     path,
		handlers: make(map[httpMethod]Endpoint),
	}
}

func (i interceptor) addMethodHandler(m httpMethod, e Endpoint) {
	if !m.isValid() {
		return
	}

	i.handlers[m] = e
}
