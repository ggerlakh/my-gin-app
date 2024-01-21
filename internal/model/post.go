package model

import "gorm.io/gorm"

// Post представляет собой модель публикации для БД.
type Post struct {
	gorm.Model
	Title   string
	Content string
	UserID  uint
}
