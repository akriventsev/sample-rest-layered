package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"
)

type JSONB map[string]interface{}

// Scan scan value into Jsonb, implements sql.Scanner interface.
func (j *JSONB) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}

	result := JSONB{}
	err := json.Unmarshal(bytes, &result)
	*j = result

	return err
}

// Value return json value, implement driver.Valuer interface.
func (j JSONB) Value() (driver.Value, error) {
	if len(j) == 0 {
		return nil, nil
	}

	v, err := json.Marshal(j)
	if err != nil {
		return nil, err
	}

	return v, nil
}

type JSONTime time.Time

// JSONTime.
func (jt *JSONTime) String() string {
	t := time.Time(*jt)

	return t.Format("20060102")
}

func (jt JSONTime) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, jt.String())), nil
}

func (jt *JSONTime) UnmarshalJSON(b []byte) error {
	timeString := strings.Trim(string(b), `"`)

	t, err := time.Parse("20060102", timeString)
	if err == nil {
		*jt = JSONTime(t)

		return nil
	}

	return fmt.Errorf("invalid date format: %s", timeString)
}

func (jt *JSONTime) ToTime() time.Time {
	return time.Time(*jt)
}
