package routes

import (
	"github.com/ODILJON-QOBILOV/microservices/ecommerce/users_service/internal/handlers"
	"github.com/ODILJON-QOBILOV/microservices/ecommerce/users_service/pkg/utils"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
    r := gin.Default()

    r.POST("/register", handlers.Register)
    r.POST("/login", handlers.Login)

    auth := r.Group("")
    auth.Use(utils.JWTAuth())
    {
        auth.GET("/profile", handlers.Profile)
        auth.POST("/product/create", handlers.CreateProduct)
    }

    return r
}
