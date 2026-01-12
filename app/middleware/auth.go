package middleware

import (
	"net/http"

	"go-admin/DB"
	systemservice "go-admin/app/service/system"
	"go-admin/common/constant"
	"go-admin/common/redis"
	"go-admin/common/token"
	"go-admin/config"
	"go-admin/response"

	"github.com/gin-gonic/gin"
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
		if err := token.ValidToken(claims); err != nil {
			token.RefreshToken(tokenStr)
		}
		ctx.Set("userId", claims.UserId)
		ctx.Set("username", claims.Username)
		ctx.Next()
	}
}

func HasPerm(perms ...string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userId := ctx.GetInt("userId")
		if userId == constant.ADMIN_ID {
			ctx.Next()
			return
		}

		if hasPerm := (&systemservice.UserService{}).UserHasPerms(userId, perms); !hasPerm {
			response.NewError(nil).SetCode(601).SetMsg("权限不足").Json(ctx)
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
