package models

import (
	"strings"
	"time"
)

type Date time.Time

const dateFormat = "2006-1-2"

func (d *Date) UnmarshalJSON(b []byte) error {
	str := strings.Trim(string(b), `"`)
	if str == "" {
		*d = Date(time.Time{})
		return nil
	}
	t, err := time.Parse(dateFormat, str)
	if err != nil {
		return err
	}
	*d = Date(t)
	return nil
}

func (d Date) MarshalJSON() ([]byte, error) {
	t := time.Time(d)
	if t.IsZero() {
		return []byte(`""`), nil
	}
	return []byte(`"` + t.Format(dateFormat) + `"`), nil
}

func (d Date) ToTime() time.Time {
	return time.Time(d)
}
