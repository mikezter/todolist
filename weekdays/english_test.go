package weekdays

import (
	"testing"
	"time"
)

func TestPivot(t *testing.T) {
	y, m, d := 2017, time.Month(12), 30
	date := time.Date(y, m, d, 15, 53, 12, 123445, time.Local)
	expected := time.Date(y, m, d, 0, 0, 0, 0, time.UTC)
	actual := pivot(date)
	if actual != expected {
		t.Error(actual)
	}
}

func TestNextMonday(t *testing.T) {
	dates := []struct{ input, actual, expected time.Time }{
		{
			input:    newDate(2017, 5, 7), // sunday
			actual:   nextMonday(newDate(2017, 5, 7)),
			expected: newDate(2017, 5, 8),
		}, {
			input:    newDate(2017, 5, 6), // saturday
			actual:   nextMonday(newDate(2017, 5, 6)),
			expected: newDate(2017, 5, 8),
		}, {}, {
			input:    newDate(2017, 5, 8), // monday
			actual:   nextMonday(newDate(2017, 5, 8)),
			expected: newDate(2017, 5, 15),
		}, {
			input:    newDate(2017, 5, 9), // tuesday
			actual:   nextMonday(newDate(2017, 5, 9)),
			expected: newDate(2017, 5, 15),
		}, {
			input:    newDate(2017, 5, 13), // saturday
			actual:   nextMonday(newDate(2017, 5, 13)),
			expected: newDate(2017, 5, 15),
		},
	}

	for _, c := range dates {
		if c.expected != c.actual {
			t.Error(c.input, c.actual, c.expected)
		}
	}
}

func newDate(y, m, d int) time.Time {
	return time.Date(y, time.Month(m), d, 0, 0, 0, 1, time.UTC)
}
