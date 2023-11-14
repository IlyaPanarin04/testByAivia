package service

import (
	"github.com/gin-gonic/gin"
	"testByAivia/internal/repository"
)

type Service struct {
	dataStorePostgres repository.DataStorePostgres
	server            *gin.Engine
}

func NewService(dataStorePostgres repository.DataStorePostgres, server *gin.Engine) *Service {
	return &Service{
		dataStorePostgres: dataStorePostgres,
		server:            server,
	}
}

func (s *Service) StartServer() {
	s.server.GET("/auth", s.Auth)

	s.server.Run()
}
