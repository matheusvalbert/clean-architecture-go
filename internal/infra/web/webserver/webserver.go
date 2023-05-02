package webserver

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type WebServer struct {
	Router        chi.Router
	Handlers      []Handler
	WebServerPort string
}

type Handler struct {
	method  string
	path    string
	handler http.HandlerFunc
}

func NewWebServer(serverPort string) *WebServer {
	return &WebServer{
		Router:        chi.NewRouter(),
		Handlers:      []Handler{},
		WebServerPort: serverPort,
	}
}

func (s *WebServer) AddMethod(method string, path string, handler http.HandlerFunc) {
	newHandler := Handler{
		method:  method,
		path:    path,
		handler: handler,
	}
	s.Handlers = append(s.Handlers, newHandler)
}

// loop through the handlers and add them to the router
// register middeleware logger
// start the server
func (s *WebServer) Start() {
	s.Router.Use(middleware.Logger)
	for _, handler := range s.Handlers {
		s.Router.Method(handler.method, handler.path, handler.handler)
	}
	http.ListenAndServe(s.WebServerPort, s.Router)
}
