package methods

import (
	"net/http"

	"github.com/eadenink/go-events/models"
	"github.com/gin-gonic/gin"
)

func SignUp(context *gin.Context) {
	var user models.User

	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = user.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": "Can't save user",
		})
		return
	}

	context.JSON(http.StatusCreated, gin.H{
		"user": user,
	})
}
