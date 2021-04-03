package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/ekyoung/gin-nice-recovery"
	"github.com/paradewisudaitb/Backend/common/serializer"
)

func InitErrorHandler(g *gin.Engine) {
	g.Use(nice.Recovery(CustomRecovery))
}

func CustomRecovery(c *gin.Context, recovered interface{}) {
	if err, ok := recovered.(string); ok {
		c.JSON(http.StatusBadRequest, serializer.ResponseBase{
			Code:    http.StatusBadRequest,
			Message: err,
		})
	}

	c.AbortWithStatus(http.StatusBadRequest)
}
