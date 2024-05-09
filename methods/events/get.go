package methods

import (
	"net/http"

	"github.com/eadenink/go-events/models"
	"github.com/gin-gonic/gin"
)

func GetEvents(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"events": models.GetEvents(),
	})
}
