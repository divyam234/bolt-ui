package service

import (
	httpPort "github.com/divyam234/bolt-ui/ports/http"
)

type Service struct {
	HTTPServer *httpPort.Server
}

func NewService(httpServer *httpPort.Server) *Service {
	return &Service{
		HTTPServer: httpServer,
	}
}
