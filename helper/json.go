package helper

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
)

func ReadFromRequestBody(c *gin.Context, result interface{}) {
	decoder := json.NewDecoder(c.Request.Body)
	err := decoder.Decode(result)
	PanicIfError(err)
}
