package httpwrapper

import (
	"fmt"
	"net/http"
)

type Interceptor struct {
	handlers map[httpMethod]Endpoint
}

func (i Interceptor) intercept(w http.ResponseWriter, r *http.Request) {
	method := httpMethod(r.Method)
	fmt.Println(method)
	endpoint, ok := i.handlers[method]
	if !ok {
		http.NotFound(w, r)
		return
	}
	endpoint.handler(w, r)
}

func newInterceptor() Interceptor {
	return Interceptor{
		handlers: make(map[httpMethod]Endpoint),
	}
}

func (i Interceptor) addMethodHandler(m httpMethod, e Endpoint) {
	if !m.isValid() {
		return
	}

	i.handlers[m] = e
}
