package controller

import (
	"errors"
	"net/http"

	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
	"github.com/paradewisudaitb/Backend/common/constant/statuscode"
	"github.com/paradewisudaitb/Backend/common/serializer"
	"github.com/paradewisudaitb/Backend/module/entity"
	uuid "github.com/satori/go.uuid"
)

type JurusanController struct {
	usecase entity.JurusanUseCase
}

func NewJurusanController(router *gin.Engine, ju entity.JurusanUseCase) JurusanController {
	cont := JurusanController{usecase: ju}
	jurusanGroup := router.Group("/jurusan")
	{
		jurusanGroup.POST("/", cont.CreateJurusan)
		jurusanGroup.PUT("/", cont.UpdateJurusan)
		jurusanGroup.DELETE("/:id", cont.DeleteJurusan)
		jurusanGroup.GET("/:id", cont.GetJurusan)
	}
	return cont
}

func (a JurusanController) CreateJurusan(ctx *gin.Context) {
	var j entity.CreateJurusanSerializer
	if err := ctx.ShouldBindJSON(&j); err != nil {
		// Error dari post
		panic(statuscode.UncompatibleJSON.String())
	}
	if err := serializer.IsValid(j); err != nil {
		panic(statuscode.UncompatibleJSON.String())
	}

	if err := a.usecase.CreateJurusan(j); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, serializer.RESPONSE_NOT_FOUND)
			return
		}
		panic(err.Error())
	}

	ctx.JSON(http.StatusOK, serializer.RESPONSE_OK)
	return
}

func (a JurusanController) UpdateJurusan(ctx *gin.Context) {
	var j entity.UpdateJurusanSerializer
	if err := ctx.ShouldBindJSON(&j); err != nil {
		panic(err.Error())
	}

	if err := a.usecase.UpdateJurusan(j); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, serializer.RESPONSE_NOT_FOUND)
			return
		}
		panic(err.Error())
	}

	ctx.JSON(http.StatusOK, serializer.RESPONSE_OK)
	return
}

func (a JurusanController) DeleteJurusan(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		panic(statuscode.EmptyParam.String())
	}

	idToUuid := uuid.FromStringOrNil(id)
	if uuid.Equal(idToUuid, uuid.Nil) {
		panic(statuscode.UnknownUUID.String())
	}

	if err := a.usecase.DeleteJurusan(idToUuid); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, serializer.RESPONSE_NOT_FOUND)
			return
		}
		panic(err.Error())
	}

	ctx.JSON(http.StatusOK, serializer.RESPONSE_OK)
	return
}

func (a JurusanController) GetJurusan(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		panic(statuscode.EmptyParam.String())
	}

	idToUuid := uuid.FromStringOrNil(id)
	if uuid.Equal(idToUuid, uuid.Nil) {
		panic(statuscode.UnknownUUID.String())
	}

	result, err := a.usecase.GetJurusan(idToUuid)
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
