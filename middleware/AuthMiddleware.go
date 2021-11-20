package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/techoc/ginessential/common"
	"github.com/techoc/ginessential/model"
	"github.com/techoc/ginessential/response"
	"net/http"
	"strings"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 	获取authorization header
		tokenString := ctx.GetHeader("Authorization")

		//验证格式
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"code": http.StatusUnauthorized,
				"msg":  "权限不足",
			})
			//抛弃这次请求
			ctx.Abort()
			return
		}

		tokenString = tokenString[7:]
		token, claims, err := common.ParseToken(tokenString)
		if err != nil || !token.Valid {
			response.Response(ctx, http.StatusUnauthorized, 401, nil, "权限不足")
			//抛弃这次请求
			ctx.Abort()
			return
		}

		//验证通过后获取claim中的userId
		userId := claims.UserId
		DB := common.GetDB()
		var user model.User
		DB.First(&user, userId)

		//用户不存在
		if user.ID == 0 {
			response.Response(ctx, http.StatusUnauthorized, 401, nil, "权限不足")
			//抛弃这次请求
			ctx.Abort()
			return
		}

		//用户存在 将user的信息写入上下文
		ctx.Set("user", user)

		ctx.Next()

	}
}
