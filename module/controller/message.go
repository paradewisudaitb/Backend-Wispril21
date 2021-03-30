package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/paradewisudaitb/Backend/common/constant/statuscode"
	"github.com/paradewisudaitb/Backend/common/serializer"
	"github.com/paradewisudaitb/Backend/module/entity"
	uuid "github.com/satori/go.uuid"
)

type MessageController struct {
	usecase entity.MessageUsecase
}

func NewMessageController(router *gin.Engine, mu entity.MessageUsecase) MessageController {
	cont := MessageController{usecase: mu}
	messageGroup := router.Group("/message")
	{
		messageGroup.POST("/", cont.CreateMessage)
		messageGroup.DELETE("/:id", cont.DeleteMessage)
		messageGroup.GET("/:id", cont.GetMessage)
	}
	return cont
}

// CreateMessage(ctx *gin.Context)
// DeleteMessage(ctx *gin.Context)
// GetMessage(ctx *gin.Context)
func (a MessageController) CreateMessage(ctx *gin.Context) {
	var j entity.CreateMessageSerializer
	if err := ctx.ShouldBindJSON(&j); err != nil {
		// Error dari post
		fmt.Println(j)
		panic(statuscode.UncompatibleJSON.String())
	}
	if err := serializer.IsValid(j); err != nil {
		panic(statuscode.UncompatibleJSON.String())
	}

	if err := a.usecase.CreateMessage(j); err != nil {
		panic(err.Error())
	}

	ctx.JSON(http.StatusOK, serializer.RESPONSE_OK)
}

func (a MessageController) DeleteMessage(ctx *gin.Context) {
	id := ctx.Param("id")

	if id == "" {
		panic(statuscode.EmptyParam.String())
	}

	idToUuid := uuid.FromStringOrNil(id)
	if uuid.Equal(idToUuid, uuid.Nil) {
		panic(statuscode.UnknownUUID.String())
	}

	if err := a.usecase.DeleteMessage(idToUuid); err != nil {
		panic(err.Error())
	}

	ctx.JSON(http.StatusOK, serializer.RESPONSE_OK)

}

func (a MessageController) GetMessage(ctx *gin.Context) {
	id := ctx.Param("id")

	if id == "" {
		panic(statuscode.EmptyParam.String())
	}

	idToUuid := uuid.FromStringOrNil(id)
	if uuid.Equal(idToUuid, uuid.Nil) {
		panic(statuscode.UnknownUUID.String())
	}

	result, err := a.usecase.GetMessage(idToUuid)
	if err != nil {
		panic(err.Error())
	}
	if len(result) == 0 {
		result = make([]entity.Message, 0)
	}

	ctx.JSON(http.StatusOK, serializer.ResponseData{
		ResponseBase: serializer.RESPONSE_OK,
		Data:         result})

}
