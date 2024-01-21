package handler

import (
	"gin-server/internal/database"
	"gin-server/internal/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Тестовая функция
func Ping(c *gin.Context) {
	value := c.Params.ByName("value")
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
		"value":   value,
	})
}

// GetUsers возвращает список всех пользователей.
func GetUsers(c *gin.Context) {
	var users []model.User
	if err := database.DB.Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": users})
}

// CreateUser создает нового пользователя.
func CreateUser(c *gin.Context) {
	// Валидация входного запроса.
	var input struct {
		Name  string `json:"name" binding:"required"`
		Email string `json:"email" binding:"required,email"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Создание пользователя.
	user := model.User{Name: input.Name, Email: input.Email}
	if err := database.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
}
