package controllers

import (
	"net/http"

	"gymdb/models"

	"gymdb/config"

	"github.com/gin-gonic/gin"
)

func Greet(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "ok",
		"message": "welcome to gymdb api",
	})
}

// getAllGyms responds with the list of all gyms as JSON.
func GetAllGyms(c *gin.Context) {
	var allGyms []models.Gym
	config.DB.Find(&allGyms)

	c.JSON(http.StatusOK, gin.H{
		"status":       "ok",
		"totalResults": len(allGyms),
		"gyms":         allGyms,
	})
}

// getGymByID locates the gym whose ID value matches the id parameter
// sent by the client, then returns that gym as a response.
func GetGymById(c *gin.Context) {
	id := c.Param("id")

	var gym models.Gym
	// result := config.DB.First(&gym, id) will result in sql injection
	result := config.DB.First(&gym, "ID = ?", id)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
		"gym":    gym,
	})
}

// addGym adds a gym from JSON received in the request body.
func AddGym(c *gin.Context) {
	var newGym models.Gym

	if err := c.BindJSON(&newGym); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	result := config.DB.Create(&newGym)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": result.Error.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"status": "ok",
		"gym":    newGym,
	})
}
