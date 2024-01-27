package model

// Post представляет собой модель публикации для БД.
type Post struct {
	//gorm.Model
	Id      uint64 `gorm:"primaryKey"`
	Title   string
	Content string
	UserID  uint
}
