package DB

import (
	"context"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/yx1126/go-admin/config"
)

type RedisContext struct {
	ctx *redis.Client
}

func (r *RedisContext) Set(key string, value interface{}, expiration time.Duration) *redis.StatusCmd {
	return r.ctx.Set(context.Background(), key, value, expiration)
}

func (r *RedisContext) Get(key string) *redis.StringCmd {
	return r.ctx.Get(context.Background(), key)
}

func (r *RedisContext) Del(key ...string) *redis.IntCmd {
	return r.ctx.Del(context.Background(), key...)
}

var Redis *RedisContext

func InitRedis() {
	Redis.ctx = redis.NewClient(&redis.Options{
		Addr:     config.Redis.Ip + ":" + strconv.Itoa(config.Redis.Port),
		Password: config.Redis.Password,
		DB:       config.Redis.Database,
	})
	_, err := Redis.ctx.Ping(context.Background()).Result()
	if err != nil {
		panic(err)
	}
}
