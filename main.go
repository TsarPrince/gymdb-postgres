package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Gym struct {
	gorm.Model

	Name        string          `gorm:"type:text"`
	Images      pq.StringArray  `gorm:"type:text[]"`
	Location    string          `gorm:"type:text"`
	Coordinates pq.Float64Array `gorm:"type:float[]"`
	Description string          `gorm:"type:text"`
	Logo        string          `gorm:"type:text"`
	Type        string          `gorm:"type:text"`
	Amenities   pq.StringArray  `gorm:"type:text[]"`
}

var db *gorm.DB
var err error

// getAllGyms responds with the list of all gyms as JSON.
func getAllGyms(c *gin.Context) {
	var allGyms []Gym
	db.Find(&allGyms)
	c.IndentedJSON(http.StatusOK, allGyms)
}

// getGymByID locates the gym whose ID value matches the id parameter
// sent by the client, then returns that gym as a response.
func getGymById(c *gin.Context) {
	id := c.Param("id")

	var gym Gym
	db.First(&gym, id)
	c.IndentedJSON(http.StatusOK, gym)
}

// addGym adds a gym from JSON received in the request body.
func addGym(c *gin.Context) {
	var newGym Gym

	// Call BindJSON to bind the received JSON to newAlbum.
	if err := c.BindJSON(&newGym); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	// result := db.Create(&newGym)					// use result.Error
	db.Create(&newGym)
	c.IndentedJSON(http.StatusCreated, gin.H{
		"message": "success",
		"data":    newGym,
	})
}

func main() {

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

	dburi := os.Getenv("DBURI")
	port := os.Getenv("PORT")
	URI := dburi

	// Opening connection to database
	db, err = gorm.Open(postgres.Open(URI), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Connected to database")
	}

	// Migrate the schema
	db.AutoMigrate(&Gym{})

	router := gin.Default()
	router.GET("/gyms", getAllGyms)
	router.GET("/gyms/:id", getGymById)
	router.POST("/gyms", addGym)

	router.Run(fmt.Sprintf("localhost:%s", port))
}
