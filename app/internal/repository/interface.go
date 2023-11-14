package repository

import "testByAivia/internal/models"

type DataStorePostgres interface {
	Auth(user models.User) (*models.User, error)
}
