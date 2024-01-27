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
	}
	// Создаем схему приложения
	err = DB.Exec("CREATE SCHEMA IF NOT EXISTS my_gin_app;").Error
	if err != nil {
		return err
	}
	// Переходим в созданную сему приложения
	err = DB.Exec("SET search_path TO my_gin_app;").Error
	if err != nil {
		return err
	}
	// Автоматическое создание таблиц на основе моделей.
	DB.AutoMigrate(&model.User{}, &model.Post{})
	return nil
}
