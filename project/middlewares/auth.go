package middlewares

import (
	"net/http"

	"example.com/m/utils"
	"github.com/gin-gonic/gin"
)

func Authentication(context *gin.Context) {
	token := context.GetHeader("Authorization")
	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "You are not authorized to view this dashboard!"})
		return
	}

	claims, err := utils.VerifyToken(token)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "token not valid"})
		return
	}

	context.Set("userId", claims)

	context.Next()

}
