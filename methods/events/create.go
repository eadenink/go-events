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
			"error": err.Error(),
		})
		return
	}

	event.DateTime = time.Now()
	event.UserID = 1

	event.Save()

	context.JSON(http.StatusCreated, gin.H{
		"event": event,
	})
}
