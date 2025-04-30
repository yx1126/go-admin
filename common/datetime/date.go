package datetime

import (
	"database/sql/driver"
	"errors"
	"time"
)

type Date struct {
	time.Time
}

func (d *Date) UnmarshalJSON(data []byte) (err error) {
	if string(data) == "null" {
		return
	}
	d.Time, err = time.Parse("2006-01-02", string(data))
	return
}

func (d *Date) MarshalJSON() ([]byte, error) {
	if d.IsZero() {
		return []byte("null"), nil
	}
	return []byte("\"" + d.Format("2006-01-02") + "\""), nil
}

// 转换为数据库值
func (d Date) Value() (driver.Value, error) {

	if d.IsZero() {
		return nil, nil
	}

	return d.Time, nil
}

// 数据库值转换为Date
func (d *Date) Scan(value interface{}) error {

	if val, ok := value.(time.Time); ok {
		*d = Date{Time: val}
		return nil
	}

	return errors.New("无法将值转换为时间戳")
}
