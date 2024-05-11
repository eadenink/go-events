package methods

import (
	"net/http"
	"time"

	"github.com/eadenink/go-events/models"
	"github.com/gin-gonic/gin"
)

func CreateEvent(context *gin.Context) {
	var event models.Event

	err := context.ShouldBindJSON(&event)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	event.DateTime = time.Now()
	event.UserID = context.GetInt64("userId")

	err = event.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Can't save event",
		})
		return
	}

	context.JSON(http.StatusCreated, gin.H{
		"event": event,
	})
}
