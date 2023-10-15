package times

import (
	"time"
)

type Time struct {
	t time.Time
}

func Now() Time {
	return Time{time.Now()}
}

func (t *Time) ISODate() string {
	return t.t.Format("2006-01-02")
}

func (t *Time) UnmarshalText(b []byte) error {
	var err error
	t.t, err = time.Parse("2006-01-02", string(b))
	return err
}
