package httpwrapper

import "net/http"

type Context struct {
	PathParams map[string]string
	*http.Request
}

func newContext(r *http.Request, paramList []pathParamRepr) (Context, error) {
	params, err := parseParms(paramList, r.URL.Path)
	return Context{
		PathParams: params,
		Request:    r,
	}, err
}

func parseParms(paramList []pathParamRepr, path string) (map[string]string, error) {

	return make(map[string]string), nil
}
