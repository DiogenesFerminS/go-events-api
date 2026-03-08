package main

import (
	"github.com/gin-gonic/gin"
	"go_event_api.com/go_api/db"
	"go_event_api.com/go_api/routes"
)

func main() {
	db.InitDB()
	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(":8080")
}
