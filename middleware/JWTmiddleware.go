package middleware

import (
	"net/http"
	"utils"

	"github.com/gin-gonic/gin"
)

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := utils.ValidateToken(c)

		if err != nil {
			c.IndentedJSON(http.StatusUnauthorized, gin.H{
				"Unauthorized": "Authentication required",
				"error":        err.Error(),
			})
			c.Abort()
			return
		}

		c.Next()
	}
}
