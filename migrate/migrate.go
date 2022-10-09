package migrate

import (
	"gymdb/config"

	"gymdb/models"
)

func MakeMigrations() {
	// Migrate the schema
	config.DB.AutoMigrate(&models.Gym{})
}
