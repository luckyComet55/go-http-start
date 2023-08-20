package httpwrapper

import (
	"fmt"
	"log"
	"net/http"
)

type Server struct {
	Port         string
	interceptors map[string]interceptor
	multiplexor  *http.ServeMux
}

func NewServer(port string) *Server {
	return &Server{
		Port:         port,
		multiplexor:  http.NewServeMux(),
		interceptors: make(map[string]interceptor),
	}
}

func (s *Server) AddRoute(endpoints ...Endpoint) *Server {
	for _, e := range endpoints {
		i, ok := s.interceptors[e.Path]
		if !ok {
			i = newInterceptor()
			s.interceptors[e.Path] = i
		}
		i.addMethodHandler(e.Method, e)
	}
	return s
}

func (s *Server) Start() {
	for path, interceptor := range s.interceptors {
		fmt.Printf("route %s\n", path)
		s.multiplexor.HandleFunc(path, interceptor.intercept)
	}
	fmt.Println("Server started")
	log.Fatal(http.ListenAndServe(s.Port, s.multiplexor))
}
