package methods

import (
	"net/http"

	"github.com/eadenink/go-events/models"
	"github.com/eadenink/go-events/utils"
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

	token, err := utils.GenerateToken(user.ID, user.Email)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal server error",
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"access_token": token,
	})
}
