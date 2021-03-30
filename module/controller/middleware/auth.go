package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/paradewisudaitb/Backend/common/serializer"
)

func Auth(c *gin.Context) {
	if values := c.Request.Header.Get("Authorization"); len(values) > 0 {
		// Cek token
		if values == "iniadalahtoken" {
			c.Next()
			return
		}
		c.AbortWithStatusJSON(http.StatusForbidden, serializer.RESPONSE_FORBIDDEN)

	}

	c.AbortWithStatusJSON(http.StatusForbidden, serializer.RESPONSE_FORBIDDEN)
}
