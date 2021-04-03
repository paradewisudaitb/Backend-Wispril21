package controller

import (
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/paradewisudaitb/Backend/common/constant/statuscode"
	"github.com/paradewisudaitb/Backend/common/serializer"
	"github.com/paradewisudaitb/Backend/module/controller/middleware"
	"github.com/paradewisudaitb/Backend/module/entity"
	"github.com/paradewisudaitb/Backend/module/usecase"
	uuid "github.com/satori/go.uuid"
	limit "github.com/yangxikun/gin-limit-by-key"
	"golang.org/x/time/rate"
	"gorm.io/gorm"
)

type MessageController struct {
	usecase entity.MessageUsecase
}

func NewMessageController(router *gin.Engine, mu entity.MessageUsecase) MessageController {
	cont := MessageController{usecase: mu}
	messageGroup := router.Group("/message")
	{
		messageGroup.POST("/", limit.NewRateLimiter(func(c *gin.Context) string {
			return c.ClientIP() // limit rate by client ip
		}, func(c *gin.Context) (*rate.Limiter, time.Duration) {
			return rate.NewLimiter(rate.Every(5*time.Minute), 3), time.Hour * 12
		}, func(c *gin.Context) {
			ForceResponse(c, http.StatusTooManyRequests, "too_many_requests")
		}), cont.CreateMessage)
		messageGroup.DELETE("/:id", middleware.Auth, cont.DeleteMessage)
		messageGroup.GET("/wisudawan/:id", cont.GetMessage)
	}
	return cont
}

func (a MessageController) CreateMessage(ctx *gin.Context) {
	var j entity.CreateMessageSerializer
	if err := ctx.ShouldBindJSON(&j); err != nil {
		// Error dari post
		ForceResponse(ctx, http.StatusBadRequest, statuscode.UncompatibleJSON.String())
		return
	}
	if err := serializer.IsValid(j); err != nil {
		ForceResponse(ctx, http.StatusBadRequest, statuscode.UncompatibleJSON.String())
		return
	}

	if err := a.usecase.CreateMessage(j); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ForceResponse(ctx, http.StatusNotFound, statuscode.NotFound.String())
			return
		}
		ForceResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, serializer.RESPONSE_OK)
	return
}

func (a MessageController) DeleteMessage(ctx *gin.Context) {
	id := ctx.Param("id")

	if id == "" {
		ForceResponse(ctx, http.StatusNotFound, statuscode.EmptyParam.String())
		return
	}

	idToUuid := uuid.FromStringOrNil(id)
	if uuid.Equal(idToUuid, uuid.Nil) {
		ForceResponse(ctx, http.StatusNotFound, statuscode.EmptyParam.String())
		return
	}

	if err := a.usecase.DeleteMessage(idToUuid); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ForceResponse(ctx, http.StatusNotFound, statuscode.NotFound.String())
			return
		}
		ForceResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, serializer.RESPONSE_OK)
	return

}

func (a MessageController) GetMessage(ctx *gin.Context) {
	id := ctx.Param("id")

	if id == "" {
		ForceResponse(ctx, http.StatusNotFound, statuscode.EmptyParam.String())
		return
	}

	idToUuid := uuid.FromStringOrNil(id)
	if uuid.Equal(idToUuid, uuid.Nil) {
		ForceResponse(ctx, http.StatusNotFound, statuscode.EmptyParam.String())
		return
	}

	result, err := a.usecase.GetMessage(idToUuid)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ForceResponse(ctx, http.StatusNotFound, statuscode.NotFound.String())
			return
		}
		ForceResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	var parsedResult []entity.GetMessageSerializer
	if len(result) == 0 {
		parsedResult = make([]entity.GetMessageSerializer, 0)
	} else {
		parsedResult = make([]entity.GetMessageSerializer, len(result))
		for i, x := range result {
			parsedResult[i] = usecase.ConvertEntityMessageToSerializer(x)
		}
	}

	ctx.JSON(http.StatusOK, serializer.ResponseData{
		ResponseBase: serializer.RESPONSE_OK,
		Data:         parsedResult})
	return

}
