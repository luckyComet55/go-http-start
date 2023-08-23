package httpwrapper

import (
	"fmt"
	"net/http"
)

type Handler func(w http.ResponseWriter, c Context)

func ErrNotFound(w http.ResponseWriter, c Context) {
	http.NotFound(w, c.Request)
}

func UserHttpErrBuilder(msg string, statusCode int) Handler {
	return func(w http.ResponseWriter, c Context) {
		w.WriteHeader(statusCode)
		fmt.Fprintf(w, "error: %s", msg)
	}
}
