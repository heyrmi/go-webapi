package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/heyrmi/go-webapi/utils"
)

func Authenticate(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")

	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not authorised to perform this action."})
		return
	}

	validToken, err := utils.VerifyToken(token)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not authorised to perform this action."})
		return
	}

	jwtData, err := utils.GetDetailsFromJWT(validToken)

	userId, ok := jwtData["userId"].(int64)
	if !ok {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "UserId not found, or invalid user id."})
		return
	}

	context.Set("userId", userId)
	context.Next()
}
