package main

import (
	"github.com/ODILJON-QOBILOV/microservices/ecommerce/users_service/internal/routes"
	"github.com/ODILJON-QOBILOV/microservices/ecommerce/users_service/pkg"
	"github.com/gin-gonic/gin"
)

func main() {
	pkg.InitDatabase()
	
	r := gin.Default()
	r.Static("/uploads", "../uploads")
	
	r = routes.SetupRouter()
	
	r.Run(":8080")
}