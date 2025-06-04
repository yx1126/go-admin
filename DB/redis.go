package DB

import (
	"context"
	"strconv"

	"github.com/redis/go-redis/v9"
	"github.com/yx1126/go-admin/config"
)

var Redis *redis.Client

func InitRedis() {
	Redis = redis.NewClient(&redis.Options{
		Addr:     config.Redis.Ip + ":" + strconv.Itoa(config.Redis.Port),
		Password: config.Redis.Password,
		DB:       config.Redis.Database,
	})
	_, err := Redis.Ping(context.Background()).Result()
	if err != nil {
		panic(err)
	}
}
