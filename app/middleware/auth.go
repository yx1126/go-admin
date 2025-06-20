package middleware

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/yx1126/go-admin/DB"
	"github.com/yx1126/go-admin/common/redis"
	"github.com/yx1126/go-admin/common/token"
	"github.com/yx1126/go-admin/config"
	"github.com/yx1126/go-admin/response"
)

// 认证
func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authorization := ctx.GetHeader(config.Token.Header)
		if authorization == "" {
			response.NewError(nil).SetCode(http.StatusUnauthorized).SetMsg("未登录").Json(ctx)
			ctx.Abort()
			return
		}
		tokenStr, err := token.ParseHeaderToken(authorization)
		if err != nil {
			response.NewError(err).SetCode(http.StatusUnauthorized).Json(ctx)
			ctx.Abort()
			return
		}
		claims, err := token.ParseToken(tokenStr)
		if err != nil {
			response.NewError(err).SetCode(http.StatusUnauthorized).Json(ctx)
			ctx.Abort()
			return
		}
		n, err := DB.Redis.Ctx.Exists(ctx.Request.Context(), redis.UserTokenKey+claims.Uuid).Result()
		if err != nil || n != 1 {
			response.NewError(err).SetCode(http.StatusUnauthorized).SetMsg("请重新登录").Json(ctx)
			ctx.Abort()
			return
		}
		if claims.ExpiresAt != nil && claims.ExpiresAt.Before(time.Now().Add(time.Minute*10)) {
			token.RefreshToken(tokenStr)
		}
		ctx.Set("userId", claims.UserId)
		ctx.Set("username", claims.Username)
		ctx.Next()
	}
}
