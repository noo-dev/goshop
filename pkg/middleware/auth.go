package middleware

import (
	"github.com/gin-gonic/gin"
	"goshop/pkg/jtoken"
	"net/http"
)

func JwtAuthMiddleware() gin.HandlerFunc {
	return JwtMiddleware(jtoken.ACCESS_TOKEN_TYPE)
}

func JwtRefreshMiddleware() gin.HandlerFunc {
	return JwtMiddleware(jtoken.REFRESH_TOKEN_TYPE)
}

func JwtMiddleware(tokenType string) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenFromRequest := c.GetHeader("Authorization")
		if tokenFromRequest == "" {
			c.JSON(http.StatusUnauthorized, nil)
			c.Abort()
			return
		}

		payload, err := jtoken.ValidateToken(tokenFromRequest)
		if err != nil || payload == nil || payload["type"] != tokenType {
			c.JSON(http.StatusUnauthorized, nil)
			c.Abort()
			return
		}

		c.Set("userId", payload["id"])
		c.Set("role", payload["role"])
		c.Next()
	}
}
