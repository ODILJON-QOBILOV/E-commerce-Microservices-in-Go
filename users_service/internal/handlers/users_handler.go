package handlers

import (
	"net/http"
	"time"

	"github.com/ODILJON-QOBILOV/microservices/ecommerce/users_service/internal/models"
	"github.com/ODILJON-QOBILOV/microservices/ecommerce/users_service/pkg"
	"github.com/ODILJON-QOBILOV/microservices/ecommerce/users_service/pkg/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Register(c *gin.Context) {
	var input struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	db_user := pkg.DB.Where("email = ?", input.Email).First(&user)
	if db_user.Error != nil {
		if db_user.Error == gorm.ErrRecordNotFound {
			user := models.User{
				Username: input.Username,
				Email:    input.Email,
				Password: utils.Password(input.Password),
			}
			if err := pkg.DB.Create(&user).Error; err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}

			// создаём JWT токен
			token, err := utils.GenerateToken(int64(user.ID), 24*time.Hour)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}

			c.JSON(http.StatusOK, gin.H{
				"message": "user created",
				"user": gin.H{
					"id":       user.ID,
					"username": user.Username,
					"email":    user.Email,
				},
				"access_token": token,
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": db_user.Error.Error()})
		return
	}

	c.JSON(http.StatusBadRequest, gin.H{"error": "user already exists"})
}

func Login(c *gin.Context) {
	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	var user models.User
	
	result := pkg.DB.Where("email = ?", input.Email).First(&user)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": result.Error.Error()})
		return
	}
	
	if !utils.CheckPassword(input.Password, user.Password) {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid credentials"})
		return
	}
	
	token, err := utils.GenerateToken(int64(user.ID), 24*time.Hour)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	
	c.JSON(http.StatusAccepted, gin.H{"access_token": token})
}

func Profile(c *gin.Context) {
	UserID, _ := c.Get("user_id")
	var user models.User
	
	if err := pkg.DB.First(&user, UserID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	
	c.JSON(200, gin.H{
		"id": user.ID,
		"username": user.Username,
		"email": user.Email,
		"CreatedAt": user.CreatedAt,
	})
}