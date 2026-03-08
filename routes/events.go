package routes

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go_event_api.com/go_api/models"
)

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
	fmt.Println(err)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"ok":    false,
			"error": "Validation failed",
		})
		return
	}

	// event.ID = 1
	event.UserID = 1
	newEvent, err := event.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"ok":    false,
			"error": "Failure to create event",
		})
		return
	}

	context.JSON(http.StatusCreated, gin.H{
		"ok":   true,
		"data": newEvent,
	})
}

func updateEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"ok":    false,
			"error": "Invalid Id",
		})
		return
	}

	_, err = models.GetEventById(eventId)

	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{
			"ok":    false,
			"error": "Event not found",
		})

		return
	}

	var updatedEvent models.Event
	err = context.ShouldBindJSON(&updatedEvent)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"ok":    false,
			"error": "Validation failed",
		})
		return
	}

	updatedEvent.ID = eventId
	err = updatedEvent.Update()

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"ok":    false,
			"error": "Update failed",
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"ok":   true,
		"data": "Event updated successfully",
	})
}

func deleteEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"ok":    false,
			"error": "Invalid Id",
		})
		return
	}

	event, err := models.GetEventById(eventId)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"ok":    false,
			"error": "Event not found",
		})
		return
	}

	err = event.Delete()

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"ok":    false,
			"error": "failed to delete the event",
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"ok":   true,
		"data": "Event deleted successfully",
	})
}
