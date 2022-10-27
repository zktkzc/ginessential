package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"tkzc.com/ginessential/response"
)

func RecoveryMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				response.Fail(context, nil, fmt.Sprint(err))
			}
		}()

		context.Next()
	}
}
