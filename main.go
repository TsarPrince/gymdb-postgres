package main

import (
	"fmt"
	"gymdb/config"
	"gymdb/initializers"
	"gymdb/migrate"
	"gymdb/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
}

func main() {
	// connect to db
	config.ConnectToDB()

	// migrate models
	migrate.MakeMigrations()

	// initialize router
	router := gin.Default()

	// cors
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	// route endpoints
	routes.Routes(router)

	router.Run(fmt.Sprintf("0.0.0.0:%s", initializers.PORT))
}
