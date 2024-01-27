package handler

import (
	"gin-server/internal/database"
	"gin-server/internal/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// GetUser возвращает конкретного пользователя по id
func GetUser(c *gin.Context) {
	var user model.User
	input_user_id := c.Params.ByName("id")
	if err := database.DB.First(&user, input_user_id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}
	if err := database.DB.Find(&user.Posts, "user_id = ?", user.Id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}

// GetUsers возвращает список всех пользователей.
func GetUsers(c *gin.Context) {
	var users []model.User
	if err := database.DB.Find(&users).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}
	for idx, _ := range users {
		if err := database.DB.Find(&users[idx].Posts, "user_id = ?", users[idx].Id).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{"users": users})
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

// DeleteUser удаляет пользователя по id.
func DeleteUser(c *gin.Context) {
	var user model.User
	input_user_id := c.Params.ByName("id")
	// Создание пользователя.
	if err := database.DB.Delete(&user, input_user_id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.String(200, "User deleted")
}
