package controller

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/paradewisudaitb/Backend/common/constant/orgztype"
	"github.com/paradewisudaitb/Backend/common/constant/statuscode"
	"github.com/paradewisudaitb/Backend/common/serializer"
	"github.com/paradewisudaitb/Backend/module/controller/middleware"
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
		orgzGroup.POST("/", middleware.Auth, cont.CreateOrgz)
		orgzGroup.PUT("/", middleware.Auth, cont.UpdateOrgz)
		orgzGroup.DELETE("/:id", middleware.Auth, cont.DeleteOrgz)
		orgzGroup.GET("/id/:id", cont.GetByID) //TODO ganti jadi 
		orgzGroup.GET("/slug/:slug", cont.GetBySlug)
		orgzGroup.GET("/all", cont.GetAll)
	}
	return cont
}

func (o OrgzController) CreateOrgz(ctx *gin.Context) {
	var j entity.CreateOrgzSerializer
	if err := ctx.ShouldBindJSON(&j); err != nil {
		ForceResponse(ctx, http.StatusBadRequest, statuscode.UncompatibleJSON.String())
		return
	}
	if err := serializer.IsValid(j); err != nil {
		ForceResponse(ctx, http.StatusBadRequest, statuscode.UncompatibleJSON.String())
		return
	}
	enum, enumErr := orgztype.GetEnum(j.Category)
	if enumErr != nil {
		ForceResponse(ctx, http.StatusBadRequest, statuscode.UnknownType.String())
		return
	}
	j.Category = enum
	if err := o.usecase.CreateOrgz(j); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ForceResponse(ctx, http.StatusNotFound, statuscode.NotFound.String())
			return
		}

		ForceResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, serializer.RESPONSE_OK)

}

func (o OrgzController) DeleteOrgz(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ForceResponse(ctx, http.StatusNotFound, statuscode.EmptyParam.String())
		return
	}

	idToUuid := uuid.FromStringOrNil(id)
	if uuid.Equal(idToUuid, uuid.Nil) {
		ForceResponse(ctx, http.StatusBadRequest, statuscode.UnknownUUID.String())
		return
	}

	if err := o.usecase.DeleteOrgz(idToUuid); err != nil {
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

func (o OrgzController) UpdateOrgz(ctx *gin.Context) {
	var j entity.UpdateOrgzSerializer
	if err := ctx.ShouldBindJSON(&j); err != nil {
		ForceResponse(ctx, http.StatusBadRequest, statuscode.UncompatibleJSON.String())
		return
	}
	if j.Category != "" {
		enum, enumErr := orgztype.GetEnum(j.Category)
		if enumErr != nil {
			ForceResponse(ctx, http.StatusBadRequest, statuscode.UnknownType.String())
			return
		}
		j.Category = enum
	}
	if err := o.usecase.UpdateOrgz(j); err != nil {
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

func (o OrgzController) GetByID(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ForceResponse(ctx, http.StatusNotFound, statuscode.EmptyParam.String())
		return
	}

	idToUuid := uuid.FromStringOrNil(id)
	if uuid.Equal(idToUuid, uuid.Nil) {
		ForceResponse(ctx, http.StatusBadRequest, statuscode.UnknownUUID.String())
		return
	}

	result, err := o.usecase.GetOrgz(idToUuid)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ForceResponse(ctx, http.StatusNotFound, statuscode.NotFound.String())
			return
		}
		ForceResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, serializer.ResponseData{
		ResponseBase: serializer.RESPONSE_OK,
		Data:         result})
	return
}

// GetBySlug(ctx *gin.Context)
// GetAll(ctx *gin.Context)
func (o OrgzController) GetBySlug(ctx *gin.Context) {
	slug := ctx.Param("slug")
	if slug == "" {
		ForceResponse(ctx, http.StatusNotFound, statuscode.EmptyParam.String())
		return
	}

	result, err := o.usecase.GetBySlug(slug)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ForceResponse(ctx, http.StatusNotFound, statuscode.NotFound.String())
			return
		}
		ForceResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, serializer.ResponseData{
		ResponseBase: serializer.RESPONSE_OK,
		Data:         result})
	return
}
func (o OrgzController) GetAll(ctx *gin.Context) {

	result, err := o.usecase.GetAll()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ForceResponse(ctx, http.StatusNotFound, statuscode.NotFound.String())
			return
		}
		ForceResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	if len(result) == 0 {
		result = make([]entity.Orgz, 0)
	}
	ctx.JSON(http.StatusOK, serializer.ResponseData{
		ResponseBase: serializer.RESPONSE_OK,
		Data:         result})
	return
}
