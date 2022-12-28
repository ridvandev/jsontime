package gojsontime

import (
	"encoding/json"
	"time"
)

type DateTime struct {
	time.Time
}

func (t DateTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.Time.Format("2006-01-02T15:04:05.000Z"))
}

func (t *DateTime) UnmarshalJSON(data []byte) error {
	if string(data) == "\"\"" {
		return nil
	}

	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}

	var err error
	t.Time, err = time.Parse("2006-01-02T15:04:05.000Z", s)
	if err != nil {
		return err
	}

	return nil
}
