package server

import (
	"fmt"
	"net/http"
)

type Server struct {
	config *ServerConfig
}

func NewServer(c *ServerConfig) *Server {

	return &Server{
		config: c,
	}

}

func (s *Server) Start() error {
	addrress := fmt.Sprintf("%v:%v", s.config.Host, s.config.Port)
	return http.ListenAndServe(addrress, nil)

}
