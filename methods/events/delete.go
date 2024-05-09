package methods

import (
	"net/http"
	"strconv"

	"github.com/eadenink/go-events/models"
	"github.com/gin-gonic/gin"
)

func DeleteEvent(context *gin.Context) {
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

	err = event.Delete()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Can't delete event",
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"event": event,
	})
}
