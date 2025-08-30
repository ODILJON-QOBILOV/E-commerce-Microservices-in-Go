package handlers

import (
	"net/http"

	"github.com/ODILJON-QOBILOV/microservices/ecommerce/users_service/internal/models"
	"github.com/ODILJON-QOBILOV/microservices/ecommerce/users_service/pkg"
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
	
	db_user := pkg.DB.First(&input, "email = ?", input.Email)
	if db_user.Error != nil {
		if db_user.Error == gorm.ErrRecordNotFound {
			user := models.User{
				Username: input.Username,
				Email: input.Email,
				Password: input.Password,
			}
			pkg.DB.Create(&user)
		}
	}
}
