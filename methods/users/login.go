package methods

import (
	"net/http"

	"github.com/eadenink/go-events/models"
	"github.com/gin-gonic/gin"
)

func Login(context *gin.Context) {
	var user models.User

	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = user.ValidateCredentials()
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid credentials",
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"access_token": "dummy_access_token",
	})
}
