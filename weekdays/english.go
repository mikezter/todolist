package weekdays

import (
	"errors"
	"time"
)

type weekdayDict map[string]time.Weekday

// English returns a Weekdayer for english day names
func English(input string, t time.Time) Weekdayer {
	w := englishWeekdayer{input: input, pivot: pivot(t)}

	w.dict = weekdayDict{
		"mon": time.Monday,
		"tue": time.Tuesday,
		"wed": time.Wednesday,
		"thu": time.Thursday,
		"fri": time.Friday,
		"sat": time.Saturday,
		"sun": time.Sunday,
	}
	return w
}

type englishWeekdayer struct {
	input string
	pivot time.Time
	dict  weekdayDict
}

func (w englishWeekdayer) Weekday() (time.Time, error) {
	if weekday, ok := w.dict[w.input]; ok {
		today := w.pivot.Weekday()
		delta := weekday - today
		date := w.pivot.AddDate(0, 0, modPositive(int(delta), 7))
		return date, nil
	}
	return time.Now(), errors.New("weekday not recognized")
}

func modPositive(i, n int) int {
	return (i%n + n) % n
}

func pivot(t time.Time) time.Time {
	y, m, d := t.Date()
	return time.Date(y, m, d, 0, 0, 0, 0, t.Location())
}
