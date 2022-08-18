package ws

import (
	"net/http"
)

type Server struct {
	address    string
	port       string
	httpServer *http.Server
}

func NewServer(address string, port string) *Server {
	return &Server{address: address, port: port}
}

func (s *Server) Run(handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:    s.address + ":" + s.port,
		Handler: handler,
	}

	return s.httpServer.ListenAndServe()
}
