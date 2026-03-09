package routes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go_event_api.com/go_api/models"
)

func registerForEvent(context *gin.Context) {
	userId := context.GetInt64("userId")
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

	err = event.Register(userId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"ok":    false,
			"error": "Register failed",
		})
		return
	}

	context.JSON(http.StatusCreated, gin.H{
		"ok":   true,
		"data": "Registered!",
	})
}

func cancelRegistration(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"ok":    false,
			"error": "Invalid Id",
		})
		return
	}

	currentEvent, err := models.GetEventById(eventId)

	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{
			"ok":    false,
			"error": "Event not found",
		})
		return
	}

	err = currentEvent.CancelRegistration(userId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"ok":    false,
			"error": "Cancel registration failed",
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"ok":   true,
		"data": "Cancelled",
	})

}
