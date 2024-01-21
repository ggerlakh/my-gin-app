package database

import (
	"gin-server/internal/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DB инстанс.
var DB *gorm.DB

// Initialize подключается к базе данных и мигрирует модели.
func Initialize(dsn string) error {
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
		//panic("Не удалось подключиться к базе данных")
	}

	// Автоматическое создание таблиц на основе моделей.
	DB.AutoMigrate(&model.User{}, &model.Post{})
	return nil
}
