package main

import (
	"fmt"
	"net/http"

	"github.com/luckyComet55/go-http-start/httpwrapper"
)

func main() {
	server := httpwrapper.NewServer("3003")
	server.AddRoute(httpwrapper.NewEndpoint(
		"/",
		"GET",
		func(w http.ResponseWriter, c httpwrapper.Context) {
			fmt.Fprintln(w, "Велкоме!")
		},
	))
	server.AddRoute(httpwrapper.NewEndpoint(
		"/home",
		"GET",
		func(w http.ResponseWriter, c httpwrapper.Context) {
			fmt.Fprintln(w, "Home!")
		},
	))
	server.AddRoute(httpwrapper.NewEndpoint(
		"/{firstName}/{lastName}",
		"GET",
		func(w http.ResponseWriter, c httpwrapper.Context) {
			fmt.Fprintln(w, "Name!")
		},
	))
	server.Start()
}
