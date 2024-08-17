package c_type

import (
	"database/sql/driver"
	"strings"
)

type Array []string

func (t *Array) Scan(value interface{}) error {
	v, _ := value.([]byte)
	if string(v) == "" {
		*t = []string{}
		return nil
	}
	*t = strings.Split(string(v), ",")
	return nil
}

func (t Array) Value() (driver.Value, error) {
	return strings.Join(t, ","), nil
}
