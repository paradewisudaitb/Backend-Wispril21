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

type OrgzController struct {
	usecase entity.OrgzUseCase
}

func NewOrgzController(router *gin.Engine, ou entity.OrgzUseCase) entity.OrgzController {
	cont := OrgzController{usecase: ou}
	orgzGroup := router.Group("/orgz")
	{
		orgzGroup.POST("/", cont.CreateOrgz)
		orgzGroup.PUT("/", cont.UpdateOrgz)
		orgzGroup.DELETE("/:id", cont.DeleteOrgz)
		orgzGroup.GET("/:id", cont.GetOrgz)
	}
	return cont
}

func (o OrgzController) CreateOrgz(ctx *gin.Context) {
	var j entity.CreateOrgzSerializer
	if err := ctx.ShouldBindJSON(&j); err != nil {
		ForceResponse(ctx, http.StatusBadRequest, statuscode.UncompatibleJSON.String())
	}
	if err := serializer.IsValid(j); err != nil {
		ForceResponse(ctx, http.StatusBadRequest, statuscode.UncompatibleJSON.String())
	}

	if err := o.usecase.CreateOrgz(j); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ForceResponse(ctx, http.StatusNotFound, statuscode.NotFound.String())
		}

		ForceResponse(ctx, http.StatusBadRequest, err.Error())
	}

	ctx.JSON(http.StatusOK, serializer.RESPONSE_OK)

}

func (o OrgzController) DeleteOrgz(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ForceResponse(ctx, http.StatusNotFound, statuscode.EmptyParam.String())
	}

	idToUuid := uuid.FromStringOrNil(id)
	if uuid.Equal(idToUuid, uuid.Nil) {
		ForceResponse(ctx, http.StatusBadRequest, statuscode.UnknownUUID.String())
	}

	if err := o.usecase.DeleteOrgz(idToUuid); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ForceResponse(ctx, http.StatusNotFound, statuscode.NotFound.String())
		}
		ForceResponse(ctx, http.StatusBadRequest, err.Error())
	}

	ctx.JSON(http.StatusOK, serializer.RESPONSE_OK)
}

func (o OrgzController) UpdateOrgz(ctx *gin.Context) {
	var j entity.UpdateOrgzSerializer
	if err := ctx.ShouldBindJSON(&j); err != nil {
		ForceResponse(ctx, http.StatusBadRequest, statuscode.UncompatibleJSON.String())
	}

	if err := o.usecase.UpdateOrgz(j); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ForceResponse(ctx, http.StatusNotFound, statuscode.NotFound.String())
			return
		}
		ForceResponse(ctx, http.StatusBadRequest, err.Error())
	}
	ctx.JSON(http.StatusOK, serializer.RESPONSE_OK)

}

func (o OrgzController) GetOrgz(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ForceResponse(ctx, http.StatusNotFound, statuscode.EmptyParam.String())
	}

	idToUuid := uuid.FromStringOrNil(id)
	if uuid.Equal(idToUuid, uuid.Nil) {
		ForceResponse(ctx, http.StatusBadRequest, statuscode.UnknownUUID.String())
	}

	result, err := o.usecase.GetOrgz(idToUuid)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ForceResponse(ctx, http.StatusNotFound, statuscode.NotFound.String())
			return
		}
		ForceResponse(ctx, http.StatusBadRequest, err.Error())
	}

	ctx.JSON(http.StatusOK, serializer.ResponseData{
		ResponseBase: serializer.RESPONSE_OK,
		Data:         result})
	return
}
