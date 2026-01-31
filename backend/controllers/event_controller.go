package controllers

import (
	"net/http"
	"time"

	"mini-eventify-backend/config"
	"mini-eventify-backend/models"

	"github.com/gin-gonic/gin"
)
func CreateEvent(c *gin.Context) {
	var input struct {
		Title       string  `json:"title"`
		Description string  `json:"description"`
		Location    string  `json:"location"`
		Date        string  `json:"date"`
		Price       float64 `json:"price"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	eventDate, _ := time.Parse("2006-01-02", input.Date)

	event := models.Event{
		Title:       input.Title,
		Description: input.Description,
		Location:    input.Location,
		Date:        eventDate,
		Price:       input.Price,
	}

	config.DB.Create(&event)

	c.JSON(http.StatusCreated, event)
}
func GetEvents(c *gin.Context) {
	var events []models.Event
	config.DB.Find(&events)

	c.JSON(http.StatusOK, events)
}
func GetEventByID(c *gin.Context) {
	id := c.Param("id")
	var event models.Event

	if err := config.DB.First(&event, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Event not found"})
		return
	}

	c.JSON(200, event)
}