package httpwrapper

import "net/http"

type Handler func(w http.ResponseWriter, c Context)
