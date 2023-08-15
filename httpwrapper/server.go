package httpwrapper

import (
	"fmt"
	"log"
	"net/http"
)

type Server struct {
	Port         string
	interceptors map[string]Interceptor
	multiplexor  *http.ServeMux
}

func NewServer(port string) *Server {
	return &Server{
		Port:         port,
		multiplexor:  http.NewServeMux(),
		interceptors: make(map[string]Interceptor),
	}
}

func (s *Server) AddRoute(e Endpoint) *Server {
	i, ok := s.interceptors[e.path]
	if !ok {
		i = newInterceptor()
		s.interceptors[e.path] = i
	}
	i.addMethodHandler(e.method, e)
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
