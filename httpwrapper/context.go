package httpwrapper

import "net/http"

type Context struct {
	PathParams map[string]string
	*http.Request
}

func newContext(pathParams map[string]string, r *http.Request) Context {
	return Context{
		PathParams: pathParams,
		Request:    r,
	}
}

func parseParms(pathParamsSample string, path string) (map[string]string, error) {
	return make(map[string]string), nil
}
