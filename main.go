package main

import (
	"fmt"
	"net/http"

	"github.com/luckyComet55/go-http-start/httpwrapper"
)

func main() {
	server := httpwrapper.NewServer("3003")
	server.AddRoute(
		httpwrapper.NewEndpoint(
			"/",
			"GET",
			nil,
			func(w http.ResponseWriter, r *http.Request) {
				fmt.Fprintln(w, "Велкоме!")
			},
		),
		httpwrapper.NewEndpoint(
			"/home",
			"GET",
			nil,
			func(w http.ResponseWriter, r *http.Request) {
				fmt.Fprintln(w, "Home!")
			},
		))
	server.Start()
}
