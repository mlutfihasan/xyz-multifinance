package helper

import (
	"github.com/gin-gonic/gin"
)

func Auth(next func(c *gin.Context), roles []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		next(c)
	}
}
