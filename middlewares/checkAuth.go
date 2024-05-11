package middlewares

import (
	"github.com/eadenink/go-events/utils"
	"github.com/gin-gonic/gin"
)

func CheckAuth(context *gin.Context) {
	token := context.GetHeader("Authorization")
	if token == "" {
		context.AbortWithStatusJSON(401, gin.H{
			"message": "Not authorized",
		})
		return
	}

	userId, err := utils.VerifyToken(token)
	if err != nil {
		context.AbortWithStatusJSON(401, gin.H{
			"message": "Not authorized",
		})
		return
	}

	context.Set("userId", userId)
	context.Next()
}
