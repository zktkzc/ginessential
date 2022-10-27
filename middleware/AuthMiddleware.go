package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"tkzc.com/ginessential/common"
	"tkzc.com/ginessential/model"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		// 获取authorization header
		tokenString := context.GetHeader("Authorization")

		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			context.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足"})
			context.Abort()
			return
		}

		tokenString = tokenString[7:]
		token, claims, err := common.ParseToken(tokenString)
		if err != nil || !token.Valid {
			context.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足"})
			context.Abort()
			return
		}

		// 验证通过后获取claims中的userId
		userId := claims.UserId
		DB := common.GetDB()
		var user model.User
		DB.First(&user, userId)

		// 验证用户是否存在
		if user.ID == 0 {
			context.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足"})
			context.Abort()
			return
		}

		// 用户存在，将user信息写入上下文
		context.Set("user", user)

		context.Next()
	}
}
