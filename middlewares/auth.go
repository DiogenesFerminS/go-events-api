package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go_event_api.com/go_api/utils"
)

func Authenticate(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")

	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"ok":    false,
			"error": "token not found",
		})
		return
	}

	userId, err := utils.VerifyToken(token)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"ok":    false,
			"error": "Invalid token, not authorized",
		})
		return
	}

	context.Set("userId", userId)
	context.Next()
}
