package utils

import (
	"context"

	"stuff/internal/consts"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

// ShouldGetTraceID return trace id or empty if not exists
func ShouldGetTraceID(c context.Context) (traceID string) {
	it := c.Value(consts.ContextTraceID)
	if it == nil {
		log.Errorln("Could not get trace id from context")
		return
	}
	var ok bool
	if traceID, ok = it.(string); !ok {
		log.Errorf("Invalid trace id value in context: %v", it)
	}
	return
}

// DefaultGinContext 从gin.Context中获取context.Context，方便上下文控制
func DefaultGinContext(ctx *gin.Context) (c context.Context) {
	it, ok := ctx.Get(consts.GinContext)
	if !ok {
		log.Warnln("context doesn't exists")
		c = context.TODO()
		return
	}

	c, ok = it.(context.Context)
	if !ok {
		log.Warnln("invalid context value type")
		c = context.TODO()
		return
	}

	return
}
