package service

import (
	"github.com/gin-gonic/gin"
	"testByAivia/internal/models"
)

func (s *Service) Auth(ctx *gin.Context) {
	user := models.User{
		UserName: "User",
		Password: "123344",
	}

	userDTO, err := s.dataStorePostgres.Auth(user)

	ctx.JSON(200, gin.H{
		"user": userDTO,
		"err":  err,
	})
}
