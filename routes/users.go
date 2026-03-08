package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go_event_api.com/go_api/models"
	"go_event_api.com/go_api/utils"
)

func signup(context *gin.Context) {
	var user models.User

	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"ok":    false,
			"error": "Validation failed",
		})
		return
	}

	newUser, err := user.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"ok":    false,
			"error": "Failure to create user",
		})
		return
	}

	context.JSON(http.StatusCreated, gin.H{
		"ok":   true,
		"date": newUser,
	})
}

func login(context *gin.Context) {
	var user models.User

	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"ok":    false,
			"error": "Validation failed",
		})
		return
	}

	err = user.ValidateCredentials()
	// fmt.Println(err)

	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{
			"ok":    false,
			"error": err.Error(),
		})
		return
	}

	token, err := utils.GenerateToken(user.Email, user.ID)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"ok":    false,
			"error": "Generate token failed",
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"ok":   true,
		"data": token,
	})
}
