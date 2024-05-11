package methods

import (
	"net/http"
	"strconv"

	"github.com/eadenink/go-events/models"
	"github.com/gin-gonic/gin"
)

func Register(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid event id",
		})
		return
	}

	event, err := models.GetEvent(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Can't fetch event",
		})
		return
	}

	err = event.Register(userId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Can't register for event",
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "You successfully registered for event",
	})
}
