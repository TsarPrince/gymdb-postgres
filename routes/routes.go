package routes

import (
	"gymdb/controllers"

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	router.GET("/", controllers.Greet)
	router.GET("/gyms", controllers.GetAllGyms)
	router.GET("/gyms/:id", controllers.GetGymById)
	router.POST("/gyms", controllers.AddGym)
}
