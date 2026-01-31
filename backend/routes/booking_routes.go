package routes

import (
	"mini-eventify-backend/controllers"
	"mini-eventify-backend/middleware"
	

	"github.com/gin-gonic/gin"
)

func BookingRoutes(r *gin.Engine) {
	auth := r.Group("/")
	auth.Use(middleware.AuthMiddleware())
	{
		auth.POST("/bookings", controllers.CreateBooking)
		auth.GET("/bookings", controllers.GetUserBookings)
		auth.DELETE("/bookings/:id", controllers.CancelBooking)
	}
}

