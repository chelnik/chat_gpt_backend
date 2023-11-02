package http_server

import (
	"net/http"
	"openai/configs"
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) Run(cfg configs.ServerConfig) error {
	s.httpServer = &http.Server{
		Addr:    ":" + cfg.Port,
		Handler: nil,
	}

	return s.httpServer.ListenAndServe()
}
