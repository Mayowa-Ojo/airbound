package customtypes

import (
	"database/sql/driver"
	"fmt"
	"time"

	"cloud.google.com/go/civil"
)

// These types are defined to implement the ValueScanner interface
type Time civil.Time
type Date civil.Date

func (t Time) String() string {
	s := fmt.Sprintf("%02d:%02d:%02d", t.Hour, t.Minute, t.Second)
	if t.Nanosecond == 0 {
		return s
	}
	return s + fmt.Sprintf(".%09d", t.Nanosecond)
}

func (t Time) Value() (driver.Value, error) {
	return t.String(), nil
}

func (t Time) Scan(value interface{}) error {
	var s string

	switch v := value.(type) {
	case nil:
		return nil
	case []byte:
		s = string(v)
	case string:
		s = v
	case time.Time:
		tt := civil.TimeOf(v)
		t.Hour = tt.Hour
		t.Minute = tt.Minute
		t.Second = tt.Second
		t.Nanosecond = tt.Hour
		return nil
	default:
		return fmt.Errorf("unexpected type %T", v)
	}

	if s != "" {
		v, err := civil.ParseTime(s)
		if err != nil {
			return err
		}

		t.Hour = v.Hour
		t.Minute = v.Minute
		t.Second = v.Second
		t.Nanosecond = v.Nanosecond
	}

	return nil
}

func (d Date) String() string {
	return fmt.Sprintf("%04d-%02d-%02d", d.Year, d.Month, d.Day)
}

func (d Date) Value() (driver.Value, error) {
	return d.String(), nil
}

func (d Date) Scan(value interface{}) error {
	var s string

	switch v := value.(type) {
	case nil:
		return nil
	case []byte:
		s = string(v)
	case string:
		s = v
	case time.Time:
		dt := civil.DateOf(v)
		d.Year = dt.Year
		d.Month = dt.Month
		d.Day = dt.Day
		return nil
	default:
		return fmt.Errorf("unexpected type %T", v)
	}

	if s != "" {
		v, err := civil.ParseDate(s)
		if err != nil {
			return err
		}

		d.Day = v.Day
		d.Month = v.Month
		d.Year = v.Year
	}

	return nil
}

type WeekDay int

const (
	Sunday WeekDay = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)
