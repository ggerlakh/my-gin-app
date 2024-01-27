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
	// User CRUD
	router.GET("/user/:id", handler.GetUser)
	router.DELETE("/user/:id", handler.DeleteUser)
	router.POST("/user", handler.CreateUser)
	// Get all Users
	router.GET("/users", handler.GetUsers)
	// Posts CRUD
	router.GET("/post/:id", handler.GetPost)
	router.DELETE("/post/:id", handler.DeletePost)
	router.POST("/post", handler.CreatePost)
	// Get all Posts
	router.GET("/posts", handler.GetPosts)

	// Запускаем сервер.
	log.Fatal(router.Run(":8080"))
}
