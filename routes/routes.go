package routes

import (
	"github.com/gin-gonic/gin"
	"phone-go/controllers"
)

func SetupRoutes(r *gin.Engine) {
	r.GET("/phones", controllers.GetPhones)
	r.GET("/phones/:id", controllers.GetPhone)
	r.POST("/phones", controllers.CreatePhone)
	r.PUT("/phones/:id", controllers.UpdatePhone)
	r.DELETE("/phones/:id", controllers.DeletePhone)
}
