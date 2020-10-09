package middleware

import (
	"context"
	"github.com/gin-gonic/gin"
	"time"
)

func ContextTimeout(t time.Duration) func(c *gin.Context) {
	return func(c *gin.Context) {
		ctx, cancle := context.WithTimeout(c.Request.Context(), t)
		defer cancle()

		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
