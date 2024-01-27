package model

// User представляет модель пользователя для БД.
type User struct {
	//gorm.Model
	Id    uint64 `gorm:"primaryKey"`
	Name  string
	Email string `gorm:"unique"`
	Posts []Post `gorm:"foreignKey:UserID"`
}
