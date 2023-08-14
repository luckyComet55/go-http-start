package server

import (
	"fmt"
	"log"
	"strings"

	"net/http"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	method := r.Method
	if method != "GET" {
		http.NotFound(w, r)
		return
	}
	fmt.Fprint(w, "Success")
}

func nameHandler(w http.ResponseWriter, r *http.Request) {
	pathParams := strings.Split(r.URL.Path, "/")
	var res string
	pathLen := len(pathParams)

	// So, it turns out that Go adds one '/' after last sub-path of handled path
	// Therefore, we have to check that particular case
	if pathLen > 3 {
		http.NotFound(w, r)
		return
	} else if pathLen == 3 && pathParams[2] != "" {
		res = fmt.Sprintf("Hello, %s", pathParams[2])
	} else {
		res = "Hello, anonym!"
	}
	fmt.Fprint(w, res)
}

func ServerInit() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", rootHandler)
	mux.HandleFunc("/name/", nameHandler)
	log.Fatal(http.ListenAndServe(":3003", mux))
}
