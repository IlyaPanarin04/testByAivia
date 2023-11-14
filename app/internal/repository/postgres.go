package repository

import (
	"gorm.io/gorm"
	"log"
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

func (r *Repo) TestConnection() {
	log.Print("Database connection")
}
