package http_server

import (
	"net/http"
	"openai/configs"
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) Run(cfg configs.ServerConfig, handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:    ":" + cfg.Port,
		Handler: handler,
	}

	return s.httpServer.ListenAndServe()
}
