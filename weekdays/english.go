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
		"mon":       time.Monday,
		"monday":    time.Monday,
		"tue":       time.Tuesday,
		"tuesday":   time.Tuesday,
		"wed":       time.Wednesday,
		"wednesday": time.Wednesday,
		"thu":       time.Thursday,
		"thursday":  time.Thursday,
		"fri":       time.Friday,
		"friday":    time.Friday,
		"sat":       time.Saturday,
		"saturday":  time.Saturday,
		"sun":       time.Sunday,
		"sunday":    time.Sunday,
		"tod":       w.pivot.Weekday(),
		"today":     w.pivot.Weekday(),
		"tom":       w.pivot.AddDate(0, 0, 1).Weekday(),
		"tomorrow":  w.pivot.AddDate(0, 0, 1).Weekday(),
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

	switch w.input {
	case "last week":
		return nextMonday(w.pivot.AddDate(0, 0, -14)), nil
	case "next week":
		return nextMonday(w.pivot), nil
	}

	return w.pivot, errors.New("weekday not recognized")
}

func modPositive(i, n int) int {
	return (i%n + n) % n
}

func nextMonday(t time.Time) time.Time {
	if t.Weekday() == time.Sunday {
		return t.AddDate(0, 0, 1)
	}
	return t.AddDate(0, 0, int(8-t.Weekday()))
}

func pivot(t time.Time) time.Time {
	y, m, d := t.Date()
	return time.Date(y, m, d, 0, 0, 0, 0, t.Location())
}
