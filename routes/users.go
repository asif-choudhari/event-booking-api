package routes

import (
	"event-booking-api/models"
	"event-booking-api/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SignUpUser(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, models.Fail{
			StatusCode: http.StatusBadGateway,
			Message:    err.Error(),
		})
		return
	}

	userId, err := services.AddUser(&user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Fail{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, models.Success{
		StatusCode: http.StatusCreated,
		Data:       userId,
	})
}

func LoginUser(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, models.Fail{
			StatusCode: http.StatusBadGateway,
			Message:    err.Error(),
		})
		return
	}

	token, err := services.LoginUser(&user)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, models.Fail{
			StatusCode: http.StatusUnauthorized,
			Message:    err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, models.Success{
		StatusCode: http.StatusCreated,
		Data:       token,
	})
}
