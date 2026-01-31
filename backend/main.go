package main

import (
	"mini-eventify-backend/config"
	"mini-eventify-backend/models"
"mini-eventify-backend/routes"

"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	 "github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	r := gin.Default()
// ✅ CORS config
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000","https://mini-eventify-frontend.vercel.app",},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))
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
