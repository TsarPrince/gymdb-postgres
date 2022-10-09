package routes

import (
	"gymdb/controllers"
	"gymdb/middlewares"

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	router.GET("/", middlewares.AuthMiddleware, controllers.Greet)
	router.GET("/gyms", middlewares.AuthMiddleware, controllers.GetAllGyms)
	router.GET("/gyms/:id", middlewares.AuthMiddleware, controllers.GetGymById)
	router.POST("/gyms", middlewares.AuthMiddleware, controllers.AddGym)
}
