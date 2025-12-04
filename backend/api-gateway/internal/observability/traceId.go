package observability

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func CorrelationIDMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		traceID := c.GetHeader("Trace-Id")
		if traceID == "" {
			traceID = uuid.New().String() // SEMPRE único
		}

		c.Set("trace_id", traceID)
		c.Writer.Header().Set("Trace-Id", traceID)

		c.Next()
	}
}
