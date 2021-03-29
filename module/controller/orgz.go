package controller

import "github.com/paradewisudaitb/Backend/module/entity"

type OrgzController struct {
	usecase entity.OrgzUseCase
}
func NewOrgzController(router *gin.Engine, ou entity.OrgzUseCase entity.OrgzController {
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

	if err := a.usecase.CreateOrgz(j); err != nil {
		panic(err.Error())
	}

	ctx.JSON(http.StatusOK, serializer.RESPONSE_OK)
	return
}
