package httpwrapper

import (
	"fmt"
	"net/http"
	"regexp"
)

type interceptor struct {
	pathValidator *regexp.Regexp
	handlers      map[httpMethod]Endpoint
}

func (i interceptor) intercept(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	fmt.Printf("%s %s ==> %v\n", r.Method, path, i.pathValidator)
	if !i.pathValidator.MatchString(path) {
		http.NotFound(w, r)
		return
	}

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
	var pathValidator = regexp.MustCompile(fmt.Sprintf("^%s/?$", path))
	fmt.Printf("%s path validator -> %v\n", path, pathValidator)
	return interceptor{
		pathValidator: pathValidator,
		handlers:      make(map[httpMethod]Endpoint),
	}
}

func (i interceptor) addMethodHandler(m httpMethod, e Endpoint) error {
	if !m.isValid() {
		return newUnsupportedMethodError(string(m))
	}

	i.handlers[m] = e
	return nil
}
