package weekdays

import (
	"testing"
	"time"
)

func TestPivot(t *testing.T) {
	y, m, d := 2017, time.Month(12), 30
	date := time.Date(y, m, d, 15, 53, 12, 123445, time.UTC)
	expected := time.Date(y, m, d, 0, 0, 0, 0, time.UTC)
	actual := pivot(date)
	if actual != expected {
		t.Error(actual)
	}
}

func TestNextMonday(t *testing.T) {

	var input, expected, actual time.Time

	input = newDate(2017, 5, 7) // sunday
	actual = nextMonday(input)
	expected = newDate(2017, 5, 8)
	if expected != actual {
		t.Error(input, actual, expected)
	}

	input = newDate(2017, 5, 8) // monday
	actual = nextMonday(input)
	expected = newDate(2017, 5, 15)
	if expected != actual {
		t.Error(input, actual, expected)
	}

	input = newDate(2017, 5, 9) // tuesday
	actual = nextMonday(input)
	expected = newDate(2017, 5, 15)
	if expected != actual {
		t.Error(input, actual, expected)
	}

	input = newDate(2017, 5, 13) // saturday
	actual = nextMonday(input)
	expected = newDate(2017, 5, 15)
	if expected != actual {
		t.Error(input, actual, expected)
	}

}

func newDate(y, m, d int) time.Time {
	return time.Date(y, time.Month(m), d, 0, 0, 0, 1, time.UTC)
}
