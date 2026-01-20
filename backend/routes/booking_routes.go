package routes

import (
	"mini-eventify-backend/controllers"
	"mini-eventify-backend/middleware"
	

	"github.com/gin-gonic/gin"
)

func BookingRoutes(r *gin.Engine) {
	protected := r.Group("/")
	protected.Use(middleware.AuthMiddleware())

	protected.POST("/bookings", controllers.CreateBooking)
	protected.GET("/bookings", controllers.GetUserBookings)
}
