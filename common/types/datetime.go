package types

import (
	"database/sql/driver"
	"errors"
	"time"
)

type Datetime struct {
	time.Time
}

func (dt *Datetime) UnmarshalJSON(data []byte) (err error) {
	if string(data) == "null" {
		return
	}
	dt.Time, err = time.Parse("2006-01-02 15:04:05", string(data))
	return
}

func (dt *Datetime) MarshalJSON() ([]byte, error) {
	if dt.IsZero() {
		return []byte("null"), nil
	}
	return []byte("\"" + dt.Format("2006-01-02 15:04:05") + "\""), nil
}

// 转换为数据库值
func (d Datetime) Value() (driver.Value, error) {

	if d.IsZero() {
		return nil, nil
	}

	return d.Time, nil
}

// 数据库值转换为Datetime
func (d *Datetime) Scan(value interface{}) error {

	if val, ok := value.(time.Time); ok {
		*d = Datetime{Time: val}
		return nil
	}

	return errors.New("无法将值转换为时间戳")
}
