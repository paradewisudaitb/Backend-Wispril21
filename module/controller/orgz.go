package controller

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/paradewisudaitb/Backend/common/constant/statuscode"
	"github.com/paradewisudaitb/Backend/common/serializer"
	"github.com/paradewisudaitb/Backend/module/entity"
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
		// Error dari post
		panic(statuscode.UncompatibleJSON.String())
	}
	if err := serializer.IsValid(j); err != nil {
		panic(statuscode.UncompatibleJSON.String())
	}

	if err := o.usecase.CreateOrgz(j); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, serializer.RESPONSE_NOT_FOUND)
			return
		}
		panic(err.Error())
	}

	ctx.JSON(http.StatusOK, serializer.RESPONSE_OK)

}

func (o OrgzController) DeleteOrgz(ctx *gin.Context) {

}

func (o OrgzController) UpdateOrgz(ctx *gin.Context) {

}

func (o OrgzController) GetOrgz(ctx *gin.Context) {

}
