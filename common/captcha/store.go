package captcha

import (
	"context"
	"time"

	"github.com/yx1126/go-admin/DB"
	"github.com/yx1126/go-admin/common/redis"
)

type Store struct{}

func (*Store) Set(id string, value string) error {
	return DB.Redis.Set(context.Background(), redis.CaptchaCodeKey+id, value, time.Minute*5).Err()
}

// Get returns stored digits for the captcha id. Clear indicates
// whether the captcha must be deleted from the store.
func (*Store) Get(id string, clear bool) string {
	captcha, err := DB.Redis.Get(context.Background(), redis.CaptchaCodeKey+id).Result()
	if err != nil {
		return ""
	}
	if clear {
		if err = DB.Redis.Del(context.Background(), redis.CaptchaCodeKey+id).Err(); err != nil {
			return ""
		}
	}
	return captcha
}

// Verify captcha's answer directly
func (s *Store) Verify(id, answer string, clear bool) bool {
	return s.Get(id, clear) == answer
}
