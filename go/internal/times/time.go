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

func Parse(dateString string) Time {
	t, err := parseDate(dateString)
	if err != nil {
		panic(err)
	}
	return Time{t}
}

func (t Time) String() string {
	return t.ISODate()
}

func (t Time) ISODate() string {
	return t.t.Format("2006-01-02")
}

func (t Time) ShortDate() string {
	return t.t.Format("January 2, 2006")
}

func (t *Time) UnmarshalText(b []byte) error {
	var err error
	t.t, err = parseDate(string(b))
	return err
}

func parseDate(dateString string) (time.Time, error) {
	return time.Parse("2006-01-02", dateString)
}

// Compare returns which time in the following expression is earlier: a.Compare(b).
// If a is earlier, -1 is returned.
// If the times are equal, 0 is returned.
// Otherwise, 1 is returned.
func (t Time) Compare(u Time) int {
	if t.t.Equal(u.t) {
		return 0
	} else if t.t.Before(u.t) {
		return -1
	} else {
		return 1
	}
}
