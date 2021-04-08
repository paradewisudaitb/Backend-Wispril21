package controller

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/paradewisudaitb/Backend/common/constant/statuscode"
	"github.com/paradewisudaitb/Backend/common/serializer"
	"github.com/paradewisudaitb/Backend/module/controller/middleware"
	"github.com/paradewisudaitb/Backend/module/entity"
	"github.com/paradewisudaitb/Backend/module/usecase"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type WisudawanController struct {
	usecase     entity.WisudawanUsecase
	viewUsecase entity.ViewUseCase
}

func NewWisudawanController(router *gin.Engine, wu entity.WisudawanUsecase, vu entity.ViewUseCase) entity.WisudawanController {
	cont := WisudawanController{usecase: wu, viewUsecase: vu}
	wisudawanGroup := router.Group("/wisudawan")
	{
		wisudawanGroup.POST("/", middleware.Auth, cont.CreateWisudawan)
		wisudawanGroup.PUT("/", middleware.Auth, cont.UpdateWisudawan)
		wisudawanGroup.DELETE("/:id", middleware.Auth, cont.DeleteWisudawan)
		wisudawanGroup.GET("/id/:id", cont.GetWisudawan)
		wisudawanGroup.GET("/org/:slug", cont.FilterWisudawanByOrgzSlug)
		wisudawanGroup.GET("/trending", cont.Trending)
	}
	return cont
}

func (a WisudawanController) Trending(ctx *gin.Context) {
	result, err := a.viewUsecase.GetTop5()
	if err != nil {
		ForceResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	var parsedResult []struct {
		Wisudawan entity.GetSimpleWisudawanSerializer
		Count     int64
	}
	if len(result) != 0 {
		parsedResult = make([]struct {
			Wisudawan entity.GetSimpleWisudawanSerializer
			Count     int64
		}, len(result))
		for i := range result {
			parsedResult[i] = struct {
				Wisudawan entity.GetSimpleWisudawanSerializer
				Count     int64
			}{Count: result[i].Count,
				Wisudawan: usecase.ConvertEntityWisudawanToSimpleSerializer(result[i].Wisudawan)}
		}
	} else {
		parsedResult = make([]struct {
			Wisudawan entity.GetSimpleWisudawanSerializer
			Count     int64
		}, 0)
	}

	ctx.JSON(http.StatusOK, serializer.ResponseData{
		ResponseBase: serializer.RESPONSE_OK,
		Data:         parsedResult})
	return
}

func (a WisudawanController) CreateWisudawan(ctx *gin.Context) {
	var j entity.CreateWisudawanSerializer
	if err := ctx.ShouldBindJSON(&j); err != nil {
		ForceResponse(ctx, http.StatusBadRequest, statuscode.UncompatibleJSON.String())
		return

	}
	if err := serializer.IsValid(j); err != nil {
		ForceResponse(ctx, http.StatusBadRequest, statuscode.UncompatibleJSON.String())
		return
	}

	if err := a.usecase.CreateWisudawan(j); err != nil {
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

func (a WisudawanController) UpdateWisudawan(ctx *gin.Context) {
	var j entity.UpdateWisudawanSerializer
	if err := ctx.ShouldBindJSON(&j); err != nil {
		ForceResponse(ctx, http.StatusBadRequest, statuscode.UncompatibleJSON.String())
		return
	}

	if err := a.usecase.UpdateWisudawan(j); err != nil {
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

func (a WisudawanController) DeleteWisudawan(ctx *gin.Context) {
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

	if err := a.usecase.DeleteWisudawan(idToUuid); err != nil {
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

func (a WisudawanController) GetWisudawan(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ForceResponse(ctx, http.StatusNotFound, statuscode.EmptyParam.String())
		return
	}

	idToUuid := uuid.FromStringOrNil(id)
	if uuid.Equal(idToUuid, uuid.Nil) {
		ForceResponse(ctx, http.StatusNotFound, statuscode.NotFound.String())
		return
	}

	result, err := a.usecase.GetWisudawan(idToUuid)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ForceResponse(ctx, http.StatusNotFound, statuscode.NotFound.String())
			return
		}
		ForceResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	parsedResult := usecase.ConvertEntityWisudawanToSerializer(result)
	viewErr := a.viewUsecase.AddView(idToUuid, ctx.ClientIP())
	if viewErr != nil {
		fmt.Println(viewErr.Error())
	}

	ctx.JSON(http.StatusOK, serializer.ResponseData{
		ResponseBase: serializer.RESPONSE_OK,
		Data:         parsedResult})
	return
}

func (a WisudawanController) FilterWisudawanByOrgzSlug(ctx *gin.Context) {
	slug := ctx.Param("slug")
	if slug == "" {
		ForceResponse(ctx, http.StatusNotFound, statuscode.EmptyParam.String())
		return
	}

	result, err := a.usecase.FilterWisudawanByOrgzSlug(slug)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ForceResponse(ctx, http.StatusNotFound, statuscode.NotFound.String())
			return
		}
		ForceResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	var parsedResult []entity.GetSimpleWisudawanSerializer
	if len(result) == 0 {
		parsedResult = make([]entity.GetSimpleWisudawanSerializer, 0)
	} else {
		parsedResult = make([]entity.GetSimpleWisudawanSerializer, len(result))
		for i, x := range result {
			parsedResult[i] = usecase.ConvertEntityWisudawanToSimpleSerializer(x)
		}
	}

	ctx.JSON(http.StatusOK, serializer.ResponseData{
		ResponseBase: serializer.RESPONSE_OK,
		Data:         parsedResult})
	return
}
