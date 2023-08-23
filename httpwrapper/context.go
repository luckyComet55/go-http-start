package httpwrapper

import "net/http"

type Context struct {
	PathParams map[string]string
	*http.Request
}

func newContext(r *http.Request) (Context, error) {
	//	params, err := parseParms(r.URL.Path, r.URL.Path)
	// if err {
	//
	// }
	return Context{
		PathParams: make(map[string]string),
		Request:    r,
	}, nil
}

func parseParms(pathParamsSample string, path string) (map[string]string, error) {
	return make(map[string]string), nil
}
