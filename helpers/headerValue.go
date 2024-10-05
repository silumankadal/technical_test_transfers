package helpers

import (
	"github.com/gin-gonic/gin"
)

// header
func GetContentType(c *gin.Context) string {
	return c.Request.Header.Get("Content-Type")
}
