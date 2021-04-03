package controller

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/paradewisudaitb/Backend/common/constant/contenttype"
	"github.com/paradewisudaitb/Backend/common/constant/statuscode"
	"github.com/paradewisudaitb/Backend/common/serializer"
	"github.com/paradewisudaitb/Backend/module/controller/middleware"
	"github.com/paradewisudaitb/Backend/module/entity"
	"github.com/paradewisudaitb/Backend/module/usecase"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type ContentController struct {
	usecase entity.ContentUseCase
}

func NewContentController(router *gin.Engine, cu entity.ContentUseCase) entity.ContentController {
	cont := ContentController{usecase: cu}
	contentGroup := router.Group("/content")
	{
		contentGroup.POST("/", middleware.Auth, cont.CreateContent)
		contentGroup.PUT("/", middleware.Auth, cont.UpdateContent)
		contentGroup.DELETE("/:id", middleware.Auth, cont.DeleteContent)
		contentGroup.GET("/id/:id", cont.GetContent)
		contentGroup.GET("/wisudawan/:id", cont.GetContentByWisudawan)
	}
	return cont
}

func (a ContentController) CreateContent(ctx *gin.Context) {
	var j entity.CreateContentSerializer
	if err := ctx.ShouldBindJSON(&j); err != nil {
		ForceResponse(ctx, http.StatusBadRequest, statuscode.UncompatibleJSON.String())
		return
	}
	if err := serializer.IsValid(j); err != nil {
		ForceResponse(ctx, http.StatusBadRequest, statuscode.UncompatibleJSON.String())
		return
	}
	enum, enumErr := contenttype.GetEnum(j.ContentType)
	if enumErr != nil {
		ForceResponse(ctx, http.StatusBadRequest, statuscode.UnknownType.String())
		return
	}
	j.ContentType = enum

	if err := a.usecase.CreateContent(j); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ForceResponse(ctx, http.StatusNotFound, statuscode.NotFound.String())
			return
		}
		ForceResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, serializer.RESPONSE_OK)
}

func (a ContentController) UpdateContent(ctx *gin.Context) {
	var j entity.UpdateContentSerializer
	if err := ctx.ShouldBindJSON(&j); err != nil {
		ForceResponse(ctx, http.StatusBadRequest, statuscode.UncompatibleJSON.String())
		return
	}

	if j.ContentType != "" {
		enum, enumErr := contenttype.GetEnum(j.ContentType)
		if enumErr != nil {
			ForceResponse(ctx, http.StatusBadRequest, statuscode.UnknownType.String())
			return
		}
		j.ContentType = enum
	}

	if err := a.usecase.UpdateContent(j); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ForceResponse(ctx, http.StatusNotFound, statuscode.NotFound.String())
			return
		}
		ForceResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, serializer.RESPONSE_OK)
}

func (a ContentController) DeleteContent(ctx *gin.Context) {
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

	if err := a.usecase.DeleteContent(idToUuid); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ForceResponse(ctx, http.StatusNotFound, statuscode.NotFound.String())
			return
		}
		ForceResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, serializer.RESPONSE_OK)
}

func (a ContentController) GetContent(ctx *gin.Context) {
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

	result, err := a.usecase.GetContent(idToUuid)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ForceResponse(ctx, http.StatusNotFound, statuscode.NotFound.String())
			return
		}
		ForceResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK,
		serializer.ResponseData{
			ResponseBase: serializer.RESPONSE_OK,
			Data:         result,
		},
	)
}

func (a ContentController) GetContentByWisudawan(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ForceResponse(ctx, http.StatusNotFound, statuscode.EmptyParam.String())
		return
	}
	u, convertErr := strconv.ParseUint(id, 10, 32)
	if convertErr != nil {
		ForceResponse(ctx, http.StatusBadRequest, convertErr.Error())
		return
	}

	result, err := a.usecase.GetByWisudawan(uint32(u))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ForceResponse(ctx, http.StatusNotFound, statuscode.NotFound.String())
			return
		}
		ForceResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	parsedResult := usecase.ConvertEntityContentsToSerializer(result)

	ctx.JSON(http.StatusOK,
		serializer.ResponseData{
			ResponseBase: serializer.RESPONSE_OK,
			Data:         parsedResult,
		},
	)
}
