package main

import (
	"gin-server/internal/routes"
)

func main() {
	// Устанавливаем соединение с базой данных и проводим миграцию.
	// if err := database.Initialize(config.GetDatabaseConfig()); err != nil {
	// 	log.Fatalf("Ошибка настройки базы данных: %v\n", err)
	// }

	// Настраиваем и запускаем Gin роутер.
	routes.RunRouter()
}
