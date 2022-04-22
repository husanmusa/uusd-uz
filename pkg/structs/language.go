package structs

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

type Language struct {
	Uz string `json:"uz"`
	Ru string `json:"ru"`
	Ki string `json:"ki"`
}

func (a Language) Value() (driver.Value, error) {
	return json.Marshal(a)
}

func (a *Language) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &a)
}
