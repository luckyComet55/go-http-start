package httpwrapper

import (
	"fmt"
	"log"
	"net/http"
)

type endpointSlice []Endpoint
type pathTable map[string]endpointSlice

type Server struct {
	Port        string
	pathMap     pathTable
	multiplexor *http.ServeMux
}

func (s *Server) rootInterceptor(w http.ResponseWriter, r *http.Request) {
	// Intercepting stuff

	path := r.URL.Path
	method := r.Method
	var handler Handler = ErrNotFound
	fmt.Printf("intercepting route %s\n", path)
	for k, v := range s.pathMap {
		fmt.Printf("now %v\n", v[0].regexpSample)
		if !v[0].regexpSample.MatchString(path) {
			continue
		}
		fmt.Printf("found path: %s -> %s\n", k, path)
		foundCorrectMethod := false
		for _, e := range v {
			if e.Method == httpMethod(method) {
				foundCorrectMethod = true
				handler = e.Handler
				break
			}
		}
		if !foundCorrectMethod {
			handler = UserHttpErrBuilder("method not allowed", 405)
		}
		break
	}
	context, err := newContext(r)
	if err != nil {
		handler = UserHttpErrBuilder("could not parse params", 400)
	}
	fmt.Printf("intercepting end\n")
	handler(w, context)
}

func NewServer(port string) *Server {
	return &Server{
		Port:        port,
		pathMap:     make(pathTable),
		multiplexor: http.NewServeMux(),
	}
}

func (s *Server) AddRoute(endpoint Endpoint) error {
	if !endpoint.Method.isValid() {
		return newUnsupportedMethodError(string(endpoint.Method))
	}
	if _, ok := s.pathMap[endpoint.Path]; !ok {
		s.pathMap[endpoint.Path] = make(endpointSlice, 0, _SUPPORTED_METHOD_NUM)
	}
	s.pathMap[endpoint.Path] = append(s.pathMap[endpoint.Path], endpoint)
	return nil
}

func (s *Server) Start() {
	s.multiplexor.HandleFunc("/", s.rootInterceptor)
	fmt.Printf("Server started at %s\n", fmt.Sprintf("http://localhost:%s", s.Port))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", s.Port), s.multiplexor))
}
