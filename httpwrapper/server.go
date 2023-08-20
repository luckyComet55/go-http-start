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
			i = newInterceptor(e.Path)
			s.interceptors[e.Path] = i
		}
		if err := i.addMethodHandler(e.Method, e); err != nil {
			panic(err)
		}
	}
	return s
}

func (s *Server) Start() {
	for path, interceptor := range s.interceptors {
		fmt.Printf("route %s\n", path)
		s.multiplexor.HandleFunc(path, interceptor.intercept)
	}
	fmt.Printf("Server started at %s\n", fmt.Sprintf("http://localhost:%s", s.Port))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", s.Port), s.multiplexor))
}
