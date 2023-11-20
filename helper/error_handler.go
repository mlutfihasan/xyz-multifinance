package helper

import (
	"fmt"
	"runtime/debug"

	"xyz-multifinance/exception"

	"github.com/gin-gonic/gin"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println("stacktrace from panic: \n" + string(debug.Stack()))
				exception.ErrorHandler(c, err, nil)
			}
		}()
		c.Next()
	}
}
