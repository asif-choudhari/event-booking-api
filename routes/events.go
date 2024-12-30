package routes

import (
	"event-booking-api/models"
	"event-booking-api/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllEvents(ctx *gin.Context) {
	result, err := services.GetAllEvents()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Fail{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, models.Success{
		StatusCode: http.StatusOK,
		Data:       result,
	})
}

func GetEvent(ctx *gin.Context) {
	eventId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.Fail{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid Event Id",
		})
		return
	}

	event, err := services.GetEvent(eventId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Fail{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, models.Success{
		StatusCode: http.StatusOK,
		Data:       event,
	})
}

func AddEvent(ctx *gin.Context) {
	var event models.Event
	if err := ctx.BindJSON(&event); err != nil {
		ctx.JSON(http.StatusBadRequest, models.Fail{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
		})
		return
	}
	event.UserId = ctx.GetInt64("userId")
	eventId, err := services.AddEvent(&event)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Fail{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, models.Success{
		StatusCode: http.StatusCreated,
		Data:       eventId,
	})
}

func UpdateEvent(ctx *gin.Context) {
	var event models.Event
	if err := ctx.BindJSON(&event); err != nil {
		ctx.JSON(http.StatusBadRequest, models.Fail{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
		})
		return
	}

	if event.Id == 0 {
		ctx.JSON(http.StatusBadRequest, models.Fail{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid Event Id",
		})
		return
	}

	result, err := services.GetEvent(event.Id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Fail{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		})
		return
	}

	if result == nil {
		ctx.JSON(http.StatusBadRequest, models.Fail{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid Event",
		})
		return
	}

	if result.UserId != ctx.GetInt64("userId") {
		ctx.JSON(http.StatusUnauthorized, models.Fail{
			StatusCode: http.StatusUnauthorized,
			Message:    "not authorized to update this event",
		})
		return
	}

	err = services.UpdateEvent(&event)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Fail{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, models.Success{
		StatusCode: http.StatusCreated,
		Data:       event,
	})
}

func DeleteEvent(ctx *gin.Context) {
	eventId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.Fail{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid Event Id",
		})
		return
	}

	result, err := services.GetEvent(eventId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Fail{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		})
		return
	}

	if result == nil {
		ctx.JSON(http.StatusBadRequest, models.Fail{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid Event Id",
		})
		return
	}

	if result.UserId != ctx.GetInt64("userId") {
		ctx.JSON(http.StatusUnauthorized, models.Fail{
			StatusCode: http.StatusUnauthorized,
			Message:    "not authorized to delete this event",
		})
		return
	}

	err = services.DeleteEvent(eventId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Fail{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusNoContent, models.Success{
		StatusCode: http.StatusNoContent,
		Data:       nil,
	})
}
