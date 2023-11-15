package service

import (
	"github.com/gin-gonic/gin"
)

type Service struct {
	/*dataStorePostgres repository.DataStorePostgres*/
	server *gin.Engine
}

func NewService( /*dataStorePostgres repository.DataStorePostgres,*/ server *gin.Engine) *Service {
	return &Service{
		//dataStorePostgres: dataStorePostgres,
		server: server,
	}
}

func (s *Service) StartServer(port string) {
	s.server.GET("/price", s.GetPrice)

	s.server.Run(port)
}
