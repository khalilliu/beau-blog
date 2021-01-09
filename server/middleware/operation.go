package middleware

import (
	"github.com/gin-gonic/gin"
)

func OperationRecord() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Next()
	}
}

