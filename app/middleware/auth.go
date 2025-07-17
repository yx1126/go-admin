package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yx1126/go-admin/DB"
	systemservice "github.com/yx1126/go-admin/app/service/system"
	"github.com/yx1126/go-admin/common/constant"
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
