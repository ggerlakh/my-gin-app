package routes

import (
	"log"

	"gin-server/internal/handler"

	"github.com/gin-gonic/gin"
)

// RunRouter настраивает и запускает роутер Gin.
func RunRouter() {
	// Настраиваем Gin роутер.
	router := gin.Default()

	// Роуты.
	// тест
	router.GET("/ping/:value", handler.Ping)
	// Users
	//router.GET("/users/:id", handler.GetUsers)
	//router.POST("/users", handler.CreateUser)
	// Posts
	//router.POST("/post", handler.CreatePost)

	// Запускаем сервер.
	log.Fatal(router.Run(":8080"))
}
