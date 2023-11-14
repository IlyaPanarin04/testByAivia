package models

import "github.com/gofrs/uuid"

type User struct {
	Id       uuid.UUID `gorm:"primaryKey;column:id"`
	UserName string    `gorm:"column:username"`
	Password string    `gorm:"column:password"`
}
