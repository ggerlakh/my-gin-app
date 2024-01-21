package handler

import (
	"gin-server/internal/database"
	"gin-server/internal/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Функции для обработки запросов для модели Post (пример создания поста)
// CreatePost создает новую публикацию.
func CreatePost(c *gin.Context) {
	// Пример валидации входных данных.
	var input struct {
		Title   string `json:"title" binding:"required"`
		Content string `json:"content" binding:"required"`
		UserID  uint   `json:"user_id" binding:"required"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Создание публикации.
	post := model.Post{Title: input.Title, Content: input.Content, UserID: input.UserID}
	if err := database.DB.Create(&post).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": post})
}
