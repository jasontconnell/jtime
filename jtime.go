package jtime

import (
	"encoding/json"
	"time"
)

type JsonTime struct {
	time.Time
}

func (jtime *JsonTime) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}

	t, err := time.Parse("2006-01-02", s)
	jtime.Time = t
	return err
}

func (jtime *JsonTime) MarshalJSON() ([]byte, error) {
	s := jtime.Format("2006-01-02")

	return json.Marshal(s)
}
