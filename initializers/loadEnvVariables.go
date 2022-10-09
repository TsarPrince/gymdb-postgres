package initializers

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

var DBURI, PORT string

func LoadEnvVariables() {
	// loading environment variables
	err := godotenv.Load()
	if err != nil {
		// not log.Fatal for production
		fmt.Println("Error loading .env file")
	}

	// host := os.Getenv("HOST")
	// user := os.Getenv("USER")
	// password := os.Getenv("PASSWORD")
	// dbname := os.Getenv("DBNAME")
	// dbport := os.Getenv("DBPORT")
	// URI := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbname, dbport)

	DBURI = os.Getenv("DBURI")
	PORT = os.Getenv("PORT")
}
