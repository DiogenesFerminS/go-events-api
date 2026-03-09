package routes

import (
	"github.com/gin-gonic/gin"
	"go_event_api.com/go_api/middlewares"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEventById)

	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)
	authenticated.POST("/events", createEvent)
	authenticated.PATCH("/events/:id", updateEvent)
	authenticated.DELETE("/events/:id", deleteEvent)
	authenticated.POST("/events/:id/registrations", registerForEvent)
	authenticated.DELETE("/events/:id/registrations", cancelRegistration)

	server.POST("/signup", signup)
	server.POST("/login", login)
}
