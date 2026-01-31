package routes

import (
	"mini-eventify-backend/controllers"

	"github.com/gin-gonic/gin"
)
func RegisterRoutes( r *gin.Engine) {
	r.POST("/events",controllers.CreateEvent)
	r.GET("/events",controllers.GetEvents)
	r.GET("/events/:id", controllers.GetEventByID)
}