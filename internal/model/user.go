package model

import "gorm.io/gorm"

// User представляет модель пользователя для БД.
type User struct {
	gorm.Model
	Id    uint64
	Name  string
	Email string `gorm:"unique"`
	Posts []Post
}
