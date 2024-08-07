package model

import (
	"fmt"
	"io"
	"time"
)

// Date wraps time.Time to add GraphQL scalar capabilities
type Date struct {
	time.Time
}

// UnmarshalGQL parses the input value
func (t *Date) UnmarshalGQL(v interface{}) error {
	switch v := v.(type) {
	case string:
		parsedTime, err := time.Parse(time.RFC3339, v)
		if err != nil {
			return fmt.Errorf("Date must be in RFC3339 format, got %v", v)
		}
		t.Time = parsedTime
		return nil
	default:
		return fmt.Errorf("Date must be a string, got %T", v)
	}
}

// MarshalGQL writes the time value to the output
func (t Date) MarshalGQL(w io.Writer) {
	w.Write([]byte(t.Time.Format(time.RFC3339)))
}
