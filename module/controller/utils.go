package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/paradewisudaitb/Backend/common/serializer"
)

func ForceResponse(ctx *gin.Context, status int, message string) {
	ctx.AbortWithStatusJSON(status,
		serializer.ResponseBase{
			Code:    status,
			Message: message,
		},
	)
}
