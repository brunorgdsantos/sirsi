package middlewares

import (
	"Sirsi/src/dtos"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ErrorMiddlewareHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		if len(c.Errors) > 0 {
			err := c.Errors.Last().Err
			apiErr, ok := err.(*dtos.APIError)
			if ok {
				c.JSON(apiErr.StatusCode, gin.H{"error": apiErr.Message})
				c.Abort()
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro interno no servidor"})
			c.Abort()
		}
	}
}
