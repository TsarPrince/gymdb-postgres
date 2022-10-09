package config

import (
	"fmt"
	"log"

	"gymdb/initializers"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	initializers.LoadEnvVariables()
}

func ConnectToDB() {
	var err error

	// Opening connection to database
	fmt.Println("Connecting to database...")
	DB, err = gorm.Open(postgres.Open(initializers.DBURI), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Connected")
	}

}
