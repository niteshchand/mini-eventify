package controllers

import (
	"net/http"

	"mini-eventify-backend/config"
	"mini-eventify-backend/models"

	"github.com/gin-gonic/gin"
)
func CreateBooking(c *gin.Context) {
	userID := c.GetUint("user_id")

	var input struct {
		EventID uint `json:"event_id"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	booking := models.Booking{
		UserID:  userID,
		EventID: input.EventID,
	}

	config.DB.Create(&booking)

	c.JSON(http.StatusCreated, booking)
}
func GetUserBookings(c *gin.Context) {
	userID := c.GetUint("user_id")

	var bookings []models.Booking
	config.DB.Where("user_id = ?", userID).Find(&bookings)

	c.JSON(http.StatusOK, bookings)
}
func CancelBooking(c *gin.Context) {
	userID := c.GetUint("userId")
	bookingID := c.Param("id")

	var booking models.Booking

	if err := config.DB.First(&booking, bookingID).Error; err != nil {
		c.JSON(404, gin.H{"error": "Booking not found"})
		return
	}

	if booking.UserID != userID {
		c.JSON(403, gin.H{"error": "Unauthorized"})
		return
	}

	config.DB.Delete(&booking)
	c.JSON(200, gin.H{"message": "Booking cancelled"})
}

