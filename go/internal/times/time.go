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

//goland:noinspection GoMixedReceiverTypes
func (t Time) String() string {
	return t.ISODate()
}

//goland:noinspection GoMixedReceiverTypes
func (t Time) ISODate() string {
	return t.t.Format("2006-01-02")
}

//goland:noinspection GoMixedReceiverTypes
func (t Time) ShortDate() string {
	return t.t.Format("January 2, 2006")
}

//goland:noinspection GoMixedReceiverTypes
func (t *Time) UnmarshalText(b []byte) error {
	var err error
	t.t, err = parseDate(string(b))
	return err
}

func parseDate(dateString string) (time.Time, error) {
	return time.Parse("2006-01-02", dateString)
}
