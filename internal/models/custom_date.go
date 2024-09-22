package models

import (
	"encoding/json"
	"strings"
	"time"
)

type CustomDate struct {
	time.Time
}

const CustomDateFormat = "02.01.2006"

func (cd *CustomDate) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	if s == "null" || s == "" {
		cd.Time = time.Time{}
		return nil
	}
	t, err := time.Parse(CustomDateFormat, s)
	if err != nil {
		return err
	}
	cd.Time = t
	return nil
}

func (cd *CustomDate) MarshalJSON() ([]byte, error) {
	if cd.Time.IsZero() {
		return json.Marshal(nil)
	}
	formatted := cd.Time.Format(CustomDateFormat)
	return json.Marshal(formatted)
}
