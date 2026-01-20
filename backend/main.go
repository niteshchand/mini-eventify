package main

import (
	"mini-eventify-backend/config"
	"mini-eventify-backend/models"
"mini-eventify-backend/routes"
	"github.com/gin-gonic/gin"
	 "github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	r := gin.Default()

	// Connect database
	config.ConnectDatabase()

	// Auto create tables
	// config.DB.AutoMigrate(&models.Event{})
	config.DB.AutoMigrate(&models.Event{}, &models.User{} , &models.Booking{})



	// Test route
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Backend + DB running 🚀",
		})
	})
	routes.AuthRoutes(r)
	routes.RegisterRoutes(r)
	routes.BookingRoutes(r)
	r.Run(":8080")
}
