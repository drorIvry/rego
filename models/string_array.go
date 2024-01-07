package models

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

type StringArray []string

func (a *StringArray) Scan(value interface{}) error {
	switch v := value.(type) {
	case []byte:
		return json.Unmarshal(v, a)
	case string:
		return json.Unmarshal([]byte(v), a)
	default:
		return fmt.Errorf("unsupported type: %T", value)
	}
}

func (a StringArray) Value() (driver.Value, error) {
	return json.Marshal(a)
}
