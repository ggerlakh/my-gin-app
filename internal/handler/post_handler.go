package handler

import (
	"gin-server/internal/database"
	"gin-server/internal/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Функции для обработки запросов для модели Post (пример создания поста)
// GetPost возвращает конкретный пост по id
func GetPost(c *gin.Context) {
	var post model.Post
	input_post_id := c.Params.ByName("id")
	if err := database.DB.First(&post, input_post_id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}
	c.JSON(http.StatusOK, post)
}

// GetPosts возвращает список всех пользователей.
func GetPosts(c *gin.Context) {
	var posts []model.Post
	if err := database.DB.Find(&posts).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{"posts": posts})
}

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

// DeletePost удалить пост по id.
func DeletePost(c *gin.Context) {
	var post model.Post
	input_post_id := c.Params.ByName("id")
	// Создание пользователя.
	if err := database.DB.Delete(&post, input_post_id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.String(200, "Post deleted")
}
