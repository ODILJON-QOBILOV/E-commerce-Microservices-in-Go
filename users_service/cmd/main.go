package main

import (
	"github.com/ODILJON-QOBILOV/microservices/ecommerce/users_service/pkg"
	"github.com/gin-gonic/gin"
)

func main() {
	pkg.InitDatabase()
	
	r := gin.Default()
	
	r.Run(":8080")
}