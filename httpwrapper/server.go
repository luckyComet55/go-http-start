package httpwrapper

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
)

type methodTable map[httpMethod]Handler
type pathTable map[*regexp.Regexp]methodTable

type Server struct {
	Port        string
	pathMap     pathTable
	multiplexor *http.ServeMux
}

func (s *Server) rootInterceptor(w http.ResponseWriter, r *http.Request) {
	// Intercepting stuff
}

func NewServer(port string) *Server {
	return &Server{
		Port:        port,
		pathMap:     make(pathTable),
		multiplexor: http.NewServeMux(),
	}
}

func (s *Server) AddRoute(endpoints ...Endpoint) *Server {
	for _, e := range endpoints {
		pathRegex := e.transformPathToRegexpStr()
		if _, ok := s.pathMap[pathRegex]; !ok {
			s.pathMap[pathRegex] = make(methodTable)
		}
		s.pathMap[pathRegex][e.Method] = e.Handler
	}
	return s
}

func (s *Server) Start() {
	s.multiplexor.HandleFunc("/", s.rootInterceptor)
	fmt.Printf("Server started at %s\n", fmt.Sprintf("http://localhost:%s", s.Port))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", s.Port), s.multiplexor))
}
