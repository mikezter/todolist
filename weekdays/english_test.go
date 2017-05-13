package weekdays

import (
	"testing"
	"time"
)

func TestWeekday(t *testing.T) {
	today := newDate(2017, 5, 9) // tuesday
	dates := map[string]time.Time{
		"tod":       newDate(2017, 5, 9),  // today
		"tom":       newDate(2017, 5, 10), // tomorrow
		"tue":       newDate(2017, 5, 9),  // tuesday
		"wed":       newDate(2017, 5, 10), // wednesday
		"thu":       newDate(2017, 5, 11), // thursday
		"fri":       newDate(2017, 5, 12), // friday
		"sat":       newDate(2017, 5, 13), // saturday
		"sun":       newDate(2017, 5, 14), // sunday
		"mon":       newDate(2017, 5, 15), // monday
		"last week": newDate(2017, 5, 1),  // monday last week
		"next week": newDate(2017, 5, 15), // monday next week
	}

	for input, expected := range dates {
		w := English(input, today)
		actual, err := w.Weekday()

		if err != nil {
			t.Error(err)
		}

		if actual != expected {
			t.Error(input, actual, expected)
		}
	}
}

func TestPivot(t *testing.T) {
	y, m, d := 2017, time.Month(12), 30
	date := time.Date(y, m, d, 15, 53, 12, 123445, time.Local)
	expected := time.Date(y, m, d, 0, 0, 0, 0, time.Local)
	actual := pivot(date)
	if actual != expected {
		t.Error(date, actual, expected)
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

func TestDelta(t *testing.T) {
	days := []struct {
		a, b time.Weekday
		i    int
	}{
		{
			a: time.Monday,
			b: time.Tuesday,
			i: 1,
		}, {
			a: time.Monday,
			b: time.Thursday,
			i: 3,
		}, {
			a: time.Sunday,
			b: time.Monday,
			i: 1,
		}, {
			a: time.Saturday,
			b: time.Thursday,
			i: 5,
		}, {
			a: time.Tuesday,
			b: time.Monday,
			i: 6,
		},
	}

	var actual, expected int
	for _, c := range days {
		actual = delta(c.a, c.b)
		expected = c.i
		if actual != expected {
			t.Error(c.a, c.b, actual, expected)
		}
	}

}

func newDate(y, m, d int) time.Time {
	return time.Date(y, time.Month(m), d, 0, 0, 0, 0, time.UTC)
}
