package httpwrapper

import (
	"fmt"
	"net/http"
	"regexp"
)

type Endpoint struct {
	Path         string
	Method       httpMethod
	Handler      func(w http.ResponseWriter, c Context)
	paramsRepr   []pathParamRepr
	regexpSample *regexp.Regexp
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

func transformPathToRegexpStr(path string) *regexp.Regexp {
	correctionChecker := regexp.MustCompile(`{[0-9]+?[A-Za-z]*?}`)
	if loc := correctionChecker.FindStringIndex(path); loc != nil {
		panic(fmt.Sprintf("incorrect path variable name %s at path %s", path[loc[0]:loc[1]], path))
	}
	// TODO: filter out variable names, which start with digits
	replacer := regexp.MustCompile(`{[A-Za-z]+?[0-9]*?}`)
	// TODO: maybe optimize this part later
	res := "^" + replacer.ReplaceAllString(path, "[0-9A-Za-z]+?") + "/?$"
	return regexp.MustCompile(res)
}

func NewEndpoint(path string, method httpMethod, handler func(w http.ResponseWriter, c Context)) Endpoint {
	if path[len(path)-1] == '/' {
		path = path[0 : len(path)-1]
	}

	return Endpoint{
		Path:         path,
		Method:       method,
		Handler:      handler,
		paramsRepr:   getPathParamRepr(path),
		regexpSample: transformPathToRegexpStr(path),
	}
}
