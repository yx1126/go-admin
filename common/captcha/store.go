package captcha

import (
	"time"

	"github.com/yx1126/go-admin/DB"
	"github.com/yx1126/go-admin/common/redis"
)

type Store struct{}

func (*Store) Set(id string, value string) error {
	return DB.Redis.Set(redis.CaptchaCodeKey+id, value, time.Minute*5).Err()
}

// Get returns stored digits for the captcha id. Clear indicates
// whether the captcha must be deleted from the store.
func (s *Store) Get(id string, clear bool) string {
	captcha, err := DB.Redis.Get(redis.CaptchaCodeKey + id).Result()
	if err != nil {
		return ""
	}
	if clear {
		if err = s.Del(id); err != nil {
			return ""
		}
	}
	return captcha
}

func (*Store) Del(id string) error {
	return DB.Redis.Del(redis.CaptchaCodeKey + id).Err()
}

// Verify captcha's answer directly
func (s *Store) Verify(id, answer string, clear bool) bool {
	return s.Get(id, clear) == answer
}
