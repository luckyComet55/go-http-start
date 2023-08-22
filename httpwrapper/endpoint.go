package httpwrapper

import (
	"fmt"
	"net/http"
	"regexp"
)

type Endpoint struct {
	Path       string
	Method     httpMethod
	Handler    func(w http.ResponseWriter, c Context)
	paramsRepr []pathParamRepr
}

type pathParamRepr struct {
	idx  int
	name string
}

func (p pathParamRepr) String() string {
	return fmt.Sprintf("%s[%d]", p.name, p.idx)
}

func getPathParamRepr(pathSample string) []pathParamRepr {
	pathParams := make([]pathParamRepr, 0, 5)
	varIdx := -1
	for i := 0; i < len(pathSample); i++ {
		if pathSample[i] == '/' {
			varIdx++
			continue
		}
		varStart, varEnd := i+1, i+1
		if pathSample[i] == '{' {
			for pathSample[varEnd] != '}' {
				varEnd++
			}
			pathParams = append(pathParams, pathParamRepr{varIdx, pathSample[varStart:varEnd]})
			i = varEnd
		}
	}
	fmt.Println(pathParams)
	return pathParams
}

func (e Endpoint) transformPathToRegexpStr() *regexp.Regexp {
	path := e.Path
	fmt.Printf("Input str %s\n", path)
	replacer := regexp.MustCompile(`{[A-Za-z]+?[0-9]*?}`)
	res := replacer.ReplaceAllString(path, "[0-9A-Za-z]+?")
	fmt.Printf("Output str %s\n", res)
	return regexp.MustCompile(res)
}

func NewEndpoint(path string, method httpMethod, handler func(w http.ResponseWriter, c Context)) Endpoint {
	if path[len(path)-1] != '/' {
		path += "/"
	}

	pathParamsRepr := getPathParamRepr(path)

	return Endpoint{
		Path:       path,
		Method:     method,
		Handler:    handler,
		paramsRepr: pathParamsRepr,
	}
}
