package handlers

import (
	"fmt"
	"net/http"

	"github.com/ODILJON-QOBILOV/microservices/ecommerce/users_service/internal/models"
	"github.com/ODILJON-QOBILOV/microservices/ecommerce/users_service/pkg"
	"github.com/gin-gonic/gin"
)

func CreateProduct(c *gin.Context) {
	var input struct {
		Title       string `form:"title"`
		Description string `form:"description"`
		Price       int    `form:"price"`
		Amount      int    `form:"amount"`
		CategoryID  int    `form:"category_id"`
	}

	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	file, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Image is required"})
		return
	}

	imagePath := fmt.Sprintf("uploads/%s", file.Filename)
	if err := c.SaveUploadedFile(file, imagePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save image"})
		return
	}

	product := models.Products{
		Title:       input.Title,
		Description: input.Description,
		Price:       input.Price,
		Amount:      input.Amount,
		Image:       imagePath,
		CategoryId: input.CategoryID,
	}

	if err := pkg.DB.Create(&product).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create product"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Id": product.ID,
		"title": product.Title,
		"description": product.Description,
		"price": product.Price,
		"amount": product.Amount,
		"image": product.Image,
		"category_id": product.CategoryId,
		"CreatedAt": product.CreatedAt,
	},)
}
