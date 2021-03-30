package controller

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/paradewisudaitb/Backend/common/constant/statuscode"
	"github.com/paradewisudaitb/Backend/common/serializer"
	"github.com/paradewisudaitb/Backend/module/entity"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type WisudawanController struct {
	usecase entity.WisudawanUsecase
}

func NewWisudawanController(router *gin.Engine, wu entity.WisudawanUsecase) entity.WisudawanController {
	cont := WisudawanController{usecase: wu}
	wisudawanGroup := router.Group("/wisudawan")
	{
		wisudawanGroup.POST("/", cont.CreateWisudawan)
		wisudawanGroup.PUT("/", cont.UpdateWisudawan)
		wisudawanGroup.DELETE("/:id", cont.DeleteWisudawan)
		wisudawanGroup.GET("/:id", cont.GetWisudawan)
	}
	return cont
}

func (a WisudawanController) CreateWisudawan(ctx *gin.Context) {
	var j entity.CreateWisudawanSerializer
	if err := ctx.ShouldBindJSON(&j); err != nil {
		// Error dari post
		panic(statuscode.UncompatibleJSON.String())
	}
	if err := serializer.IsValid(j); err != nil {
		panic(statuscode.UncompatibleJSON.String())
	}

	if err := a.usecase.CreateWisudawan(j); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, serializer.RESPONSE_NOT_FOUND)
			return
		}
		panic(err.Error())
	}

	ctx.JSON(http.StatusOK, serializer.RESPONSE_OK)
	return
}

func (a WisudawanController) UpdateWisudawan(ctx *gin.Context) {
	var j entity.UpdateWisudawanSerializer
	if err := ctx.ShouldBindJSON(&j); err != nil {
		panic(err.Error())
	}

	if err := a.usecase.UpdateWisudawan(j); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, serializer.RESPONSE_NOT_FOUND)
			return
		}
		panic(err.Error())
	}

	ctx.JSON(http.StatusOK, serializer.RESPONSE_OK)
	return
}

func (a WisudawanController) DeleteWisudawan(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		panic(statuscode.EmptyParam.String())
	}

	idToUuid := uuid.FromStringOrNil(id)
	if uuid.Equal(idToUuid, uuid.Nil) {
		panic(statuscode.UnknownUUID.String())
	}

	if err := a.usecase.DeleteWisudawan(idToUuid); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, serializer.RESPONSE_NOT_FOUND)
			return
		}
		panic(err.Error())
	}

	ctx.JSON(http.StatusOK, serializer.RESPONSE_OK)
	return
}

func (a WisudawanController) GetWisudawan(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		panic(statuscode.EmptyParam.String())
	}

	idToUuid := uuid.FromStringOrNil(id)
	if uuid.Equal(idToUuid, uuid.Nil) {
		panic(statuscode.UnknownUUID.String())
	}

	result, err := a.usecase.GetWisudawan(idToUuid)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, serializer.RESPONSE_NOT_FOUND)
			return
		}
		panic(err.Error())
	}

	ctx.JSON(http.StatusOK, serializer.ResponseData{
		ResponseBase: serializer.RESPONSE_OK,
		Data:         result})
	return
}
