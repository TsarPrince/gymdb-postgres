package main

import (
	"fmt"
	"gymdb/config"
	"gymdb/initializers"
	"gymdb/routes"

	"gymdb/migrate"

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

	// route endpoints
	routes.Routes(router)

	router.Run(fmt.Sprintf("0.0.0.0:%s", initializers.PORT))
}
