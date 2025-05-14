package types

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"strconv"
)

// id list
type Long struct {
	Val int
}

func (l Long) MarshalJSON() ([]byte, error) {
	return json.Marshal(strconv.Itoa(l.Val))
}

func (l *Long) UnmarshalJSON(data []byte) error {
	var raw interface{}
	if err := json.Unmarshal(data, &raw); err != nil {
		return fmt.Errorf("invalid JSON: %w", err)
	}

	switch v := raw.(type) {
	case string:
		i, err := strconv.Atoi(v)
		if err != nil {
			return fmt.Errorf("invalid string for Long: %w", err)
		}
		l.Val = i
		return nil
	default:
		return fmt.Errorf("unsupported type %T for Long", v)
	}
}

// 转换为数据库值
func (l Long) Value() (driver.Value, error) {
	return l.MarshalJSON()
}

// 数据库值转换为Long
func (l *Long) Scan(value interface{}) error {
	if value == nil {
		return nil
	}
	switch v := value.(type) {
	case int:
		l.Val = v
	case int8:
		l.Val = int(v)
	case int32:
		l.Val = int(v)
	case int64:
		l.Val = int(v)
	default:
		return fmt.Errorf("cannot scan type %T into Longs", value)
	}
	return nil
}
