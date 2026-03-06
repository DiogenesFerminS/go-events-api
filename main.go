package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go_event_api.com/go_api/db"
	"go_event_api.com/go_api/models"
)

func main() {
	db.InitDB()
	server := gin.Default()

	server.GET("/events", getEvents)
	server.GET("/events/:id", getEventById)
	server.POST("/events", createEvent)

	server.Run(":8080")
}

func getEventById(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"ok":    false,
			"error": "Invalid Id format",
		})
		return
	}

	event, err := models.GetEventById(eventId)

	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{
			"ok":    false,
			"error": "Event Not Found",
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"ok":   false,
		"data": event,
	})
}

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"ok":    false,
			"error": "Could not fetch events. Try again later.",
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"ok":   true,
		"data": events,
	})
}

func createEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"ok":    false,
			"error": "Validation failed",
		})
		return
	}

	event.ID = 1
	event.UserID = 1
	err = event.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"ok":    false,
			"error": "Could not fetch events. Try again later.",
		})
		return
	}

	context.JSON(http.StatusCreated, gin.H{
		"ok":   true,
		"data": event,
	})
}
