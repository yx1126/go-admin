package DB

import (
	"context"
	"strconv"
	"time"

	"go-admin/config"

	"github.com/redis/go-redis/v9"
)

type RedisContext struct {
	Ctx *redis.Client
}

func (r *RedisContext) Set(key string, value interface{}, expiration time.Duration) *redis.StatusCmd {
	return r.Ctx.Set(context.Background(), key, value, expiration)
}

func (r *RedisContext) Get(key string) *redis.StringCmd {
	return r.Ctx.Get(context.Background(), key)
}

func (r *RedisContext) Del(key ...string) *redis.IntCmd {
	return r.Ctx.Del(context.Background(), key...)
}

func (r *RedisContext) HSet(key string, values ...interface{}) *redis.IntCmd {
	return r.Ctx.HSet(context.Background(), key, values...)
}

func (r *RedisContext) HGet(key, field string) *redis.StringCmd {
	return r.Ctx.HGet(context.Background(), key, field)
}

func (r *RedisContext) HDel(key string, fields ...string) *redis.IntCmd {
	return r.Ctx.HDel(context.Background(), key, fields...)
}

func (r *RedisContext) Exists(key string) bool {
	n, err := r.Ctx.Exists(context.Background(), key).Result()
	if err != nil {
		return false
	}
	if n == 1 {
		return true
	} else {
		return false
	}
}

var Redis *RedisContext

func InitRedis() {
	Redis = &RedisContext{
		Ctx: redis.NewClient(&redis.Options{
			Addr:     config.Redis.Ip + ":" + strconv.Itoa(config.Redis.Port),
			Password: config.Redis.Password,
			DB:       config.Redis.Database,
		}),
	}
	_, err := Redis.Ctx.Ping(context.Background()).Result()
	if err != nil {
		panic(err)
	}
}
