package controller

import (
	"errors"
	"net/http"

	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
	"github.com/paradewisudaitb/Backend/common/constant/statuscode"
	"github.com/paradewisudaitb/Backend/common/serializer"
	"github.com/paradewisudaitb/Backend/module/controller/middleware"
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
		jurusanGroup.POST("/", middleware.Auth, cont.CreateJurusan)
		jurusanGroup.PUT("/", middleware.Auth, cont.UpdateJurusan)
		jurusanGroup.DELETE("/:id", middleware.Auth, cont.DeleteJurusan)
		jurusanGroup.GET("/:id", cont.GetJurusan)
	}
	return cont
}

func (a JurusanController) CreateJurusan(ctx *gin.Context) {
	var j entity.CreateJurusanSerializer
	if err := ctx.ShouldBindJSON(&j); err != nil {
		ForceResponse(ctx, http.StatusBadRequest, statuscode.UncompatibleJSON.String())
	}
	if err := serializer.IsValid(j); err != nil {
		ForceResponse(ctx, http.StatusBadRequest, statuscode.UncompatibleJSON.String())
	}

	if err := a.usecase.CreateJurusan(j); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ForceResponse(ctx, http.StatusNotFound, statuscode.NotFound.String())
		}
		ForceResponse(ctx, http.StatusBadRequest, err.Error())
	}

	ctx.JSON(http.StatusOK, serializer.RESPONSE_OK)
}

func (a JurusanController) UpdateJurusan(ctx *gin.Context) {
	var j entity.UpdateJurusanSerializer
	if err := ctx.ShouldBindJSON(&j); err != nil {
		ForceResponse(ctx, http.StatusBadRequest, statuscode.UncompatibleJSON.String())
	}

	if err := a.usecase.UpdateJurusan(j); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ForceResponse(ctx, http.StatusNotFound, statuscode.NotFound.String())
		}
		ForceResponse(ctx, http.StatusBadRequest, err.Error())
	}

	ctx.JSON(http.StatusOK, serializer.RESPONSE_OK)
}

func (a JurusanController) DeleteJurusan(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ForceResponse(ctx, http.StatusNotFound, statuscode.EmptyParam.String())
	}

	idToUuid := uuid.FromStringOrNil(id)
	if uuid.Equal(idToUuid, uuid.Nil) {
		ForceResponse(ctx, http.StatusNotFound, statuscode.EmptyParam.String())
	}

	if err := a.usecase.DeleteJurusan(idToUuid); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ForceResponse(ctx, http.StatusNotFound, statuscode.NotFound.String())
		}
		ForceResponse(ctx, http.StatusBadRequest, err.Error())
	}

	ctx.JSON(http.StatusOK, serializer.RESPONSE_OK)
}

func (a JurusanController) GetJurusan(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ForceResponse(ctx, http.StatusNotFound, statuscode.EmptyParam.String())
	}

	idToUuid := uuid.FromStringOrNil(id)
	if uuid.Equal(idToUuid, uuid.Nil) {
		ForceResponse(ctx, http.StatusNotFound, statuscode.EmptyParam.String())
	}

	result, err := a.usecase.GetJurusan(idToUuid)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ForceResponse(ctx, http.StatusNotFound, statuscode.NotFound.String())
		}
		ForceResponse(ctx, http.StatusBadRequest, err.Error())
	}

	ctx.JSON(http.StatusOK,
		serializer.ResponseData{
			ResponseBase: serializer.RESPONSE_OK,
			Data:         result,
		},
	)
}
