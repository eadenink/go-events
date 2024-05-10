package middlewares

import "github.com/gin-gonic/gin"

func CheckAuth(context *gin.Context) {
	token := context.GetHeader("Authorization")
	if token == "" {
		context.JSON(401, gin.H{
			"message": "Not authorized",
		})
		context.Abort()
		return
	}

	context.Next()
}
