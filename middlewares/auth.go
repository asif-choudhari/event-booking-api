package middlewares

import (
	"event-booking-api/models"
	"event-booking-api/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Authenticate(ctx *gin.Context) {
	token := ctx.Request.Header.Get("Authorization")

	if token == "" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, models.Fail{
			StatusCode: http.StatusUnauthorized,
			Message:    "no auth token provided",
		})
		return
	}

	userId, err := utils.VerifyToken(token)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, models.Fail{
			StatusCode: http.StatusUnauthorized,
			Message:    err.Error(),
		})
		return
	}

	ctx.Set("userId", userId)
	ctx.Next()
}
