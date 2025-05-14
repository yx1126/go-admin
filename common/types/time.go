package types

import (
	"database/sql/driver"
	"errors"
	"time"
)

type Time struct {
	time.Time
}

func (t *Time) UnmarshalJSON(data []byte) (err error) {
	if string(data) == "null" {
		return
	}
	t.Time, err = time.Parse("15:04:05", string(data))
	return
}

func (t *Time) MarshalJSON() ([]byte, error) {
	if t.IsZero() {
		return []byte("null"), nil
	}
	return []byte("\"" + t.Format("15:04:05") + "\""), nil
}

// 转换为数据库值
func (t Time) Value() (driver.Value, error) {

	if t.IsZero() {
		return nil, nil
	}

	return t.Time, nil
}

// 数据库值转换为Time
func (t *Time) Scan(value interface{}) error {

	if val, ok := value.(time.Time); ok {
		*t = Time{Time: val}
		return nil
	}

	return errors.New("无法将值转换为时间戳")
}
