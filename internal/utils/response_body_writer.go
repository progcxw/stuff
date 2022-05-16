package utils

import (
	"bytes"

	"github.com/gin-gonic/gin"
)

type RespBodyWriter struct {
	gin.ResponseWriter
	Body *bytes.Buffer
}

func (w RespBodyWriter) Write(b []byte) (int, error) {
	w.Body.Write(b)
	return w.ResponseWriter.Write(b)
}

func (w RespBodyWriter) WriteString(s string) (int, error) {
	w.Body.WriteString(s)
	return w.ResponseWriter.WriteString(s)
}
