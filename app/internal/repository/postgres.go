package repository

import (
	"gorm.io/gorm"
	"log"
	"testByAivia/internal/models"
)

type Repo struct {
	conn *gorm.DB
}

func NewConnection(conn *gorm.Dialector) (*Repo, error) {
	db, err := gorm.Open(*conn, &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return &Repo{conn: db}, nil
}

func (r *Repo) Auth(user models.User) (*models.User, error) {
	tx := r.conn.Table("users").Create(&user)
	if tx.Error != nil {
		log.Fatal(tx.Error)
	}
	return &user, nil
}
