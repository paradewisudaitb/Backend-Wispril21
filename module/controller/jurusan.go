package controller

import (
	"net/http"

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
func (a *JurusanController) CreateJurusan(ctx *gin.Context) {
	var j entity.CreateJurusanSerializer
	if err := ctx.ShouldBindJSON(&j); err != nil {
		// Error dari post
		ctx.JSON(http.StatusBadRequest, serializer.ResponseBase{
			Code:    http.StatusBadRequest,
			Message: statuscode.UncompatibleJSON.String(),
		})
		return
	}

	if err := a.usecase.CreateJurusan(j); err != nil {
		ctx.JSON(http.StatusBadRequest, serializer.ResponseBase{
			Code:    http.StatusBadRequest,
			Message: statuscode.UnknownError.String(),
		})
	}

	ctx.JSON(http.StatusOK, serializer.ResponseBase{
		Code:    http.StatusOK,
		Message: statuscode.OK.String(),
	})
}

func (a *JurusanController) UpdateJurusan(ctx *gin.Context) {
	var j entity.UpdateJurusanSerializer
	if err := ctx.ShouldBindJSON(&j); err != nil {
		// Error dari post
		ctx.JSON(http.StatusBadRequest, serializer.ResponseBase{
			Code:    http.StatusBadRequest,
			Message: statuscode.UncompatibleJSON.String(),
		})
		return
	}

	if err := a.usecase.UpdateJurusan(j); err != nil {
		ctx.JSON(http.StatusBadRequest, serializer.ResponseBase{
			Code:    http.StatusBadRequest,
			Message: statuscode.UnknownError.String(),
		})
	}

	ctx.JSON(http.StatusOK, serializer.ResponseBase{
		Code:    http.StatusOK,
		Message: statuscode.OK.String(),
	})
}

func (a *JurusanController) DeleteJurusan(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, serializer.ResponseBase{
			Code:    http.StatusBadRequest,
			Message: statuscode.EmptyParam.String(),
		})
	}

	idToUuid := uuid.FromStringOrNil(id)
	if uuid.Equal(idToUuid, uuid.Nil) {
		ctx.JSON(http.StatusBadRequest, serializer.ResponseBase{
			Code:    http.StatusBadRequest,
			Message: statuscode.UnknownUUID.String(),
		})
	}

	if err := a.usecase.DeleteJurusan(idToUuid); err != nil {
		ctx.JSON(http.StatusBadRequest, serializer.ResponseBase{
			Code:    http.StatusBadRequest,
			Message: statuscode.UnknownError.String(),
		})
	}

	ctx.JSON(http.StatusOK, serializer.ResponseBase{
		Code:    http.StatusOK,
		Message: statuscode.OK.String(),
	})
}

func (a *JurusanController) GetJurusan(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, serializer.ResponseBase{
			Code:    http.StatusBadRequest,
			Message: statuscode.EmptyParam.String(),
		})
	}

	idToUuid := uuid.FromStringOrNil(id)
	if uuid.Equal(idToUuid, uuid.Nil) {
		ctx.JSON(http.StatusBadRequest, serializer.ResponseBase{
			Code:    http.StatusBadRequest,
			Message: statuscode.UnknownUUID.String(),
		})
	}

	result, err := a.usecase.GetJurusan(idToUuid)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, serializer.ResponseBase{
			Code:    http.StatusBadRequest,
			Message: statuscode.UnknownError.String(),
		})
	}

	ctx.JSON(http.StatusOK, result)
}
