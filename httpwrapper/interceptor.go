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
	endpoint, ok := i.handlers[method]
	if !ok {
		http.NotFound(w, r)
	} else {
		params, err := parseParms(endpoint.Path, path)
		if err != nil {
			http.NotFound(w, r)
			return
		}
		c := newContext(params, r)
		endpoint.Handler(w, c)
	}
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
