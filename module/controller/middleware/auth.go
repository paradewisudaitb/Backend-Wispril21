package middleware

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/paradewisudaitb/Backend/common/serializer"
)

func Auth(c *gin.Context) {
	if values := c.Request.Header.Get("Authorization"); len(values) > 0 {
		// Cek token
		token := os.Getenv("AUTH_TOKEN")
		if values == token {
			c.Next()
			return
		}
		c.AbortWithStatusJSON(http.StatusForbidden, serializer.RESPONSE_FORBIDDEN)

	}

	c.AbortWithStatusJSON(http.StatusForbidden, serializer.RESPONSE_FORBIDDEN)
}

func ResetAuth(c *gin.Context) {
	if values := c.Request.Header.Get("Authorization"); len(values) > 0 {
		// Cek token
		token := os.Getenv("RESET_TOKEN")
		if values == token {
			c.Next()
			return
		}
		c.AbortWithStatusJSON(http.StatusForbidden, serializer.RESPONSE_FORBIDDEN)

	}

	c.AbortWithStatusJSON(http.StatusForbidden, serializer.RESPONSE_FORBIDDEN)
}
