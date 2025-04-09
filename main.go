package main

import (
	"github.com/gin-gonic/gin"
	"phone-go/database"
	"phone-go/routes"
)

func main() {
	r := gin.Default()
	database.Connect()
	routes.SetupRoutes(r)
	r.Run(":8080") // API будет доступен по http://localhost:8080
}
