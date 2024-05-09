package methods

import (
	"net/http"

	"github.com/eadenink/go-events/models"
	"github.com/gin-gonic/gin"
)

func GetEvents(context *gin.Context) {
	events, err := models.GetEvents()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Can't fetch events",
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"events": events,
	})
}
