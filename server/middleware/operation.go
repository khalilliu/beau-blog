package middleware

import (
	"bytes"
	"github.com/gin-gonic/gin"
)

func OperationRecord() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Next()
	}
}

type responseBodyWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (r *responseBodyWriter) Write(b []byte) (int, error) {
	r.body.Write(b)
	return r.ResponseWriter.Write(b)
}
