package main

import (
	"gin-server/internal/config"
	"gin-server/internal/database"
	"gin-server/internal/routes"
	"log"
)

func main() {
	// Устанавливаем соединение с базой данных и создаем таблицы.
	if err := database.Initialize(config.GetDatabaseConfig()); err != nil {
		log.Fatalf("Ошибка настройки базы данных: %v\n", err)
	}

	// Настраиваем и запускаем Gin роутер.
	routes.RunRouter()
}
