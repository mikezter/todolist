package weekdays

import "time"

// Weekdayer has Weekday() method
type Weekdayer interface {
	Weekday() (time.Time, error)
}
