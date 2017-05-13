package todolist

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/gammons/todolist/weekdays"
)

type Parser struct {
	input string
}

func (p Parser) ParseNewTodo() *Todo {
	todo := NewTodo()
	todo.Subject = p.Subject()
	todo.Projects = p.Projects()
	todo.Contexts = p.Contexts()
	if p.hasDue() {
		todo.Due = p.Due(time.Now())
	}
	return todo
}

func (p Parser) parseId() int {
	r := regexp.MustCompile(`(\d+)`)
	matches := r.FindStringSubmatch(p.input)
	if len(matches) == 0 {
		fmt.Println("Could not match id")
		return -1
	}
	id, err := strconv.Atoi(matches[1])
	if err != nil {
		fmt.Println("Invalid id.")
		return -1
	}
	return id
}

func (p Parser) parseSubject() string {
	r := regexp.MustCompile(`(\d+) (.*)`)
	matches := r.FindStringSubmatch(p.input)
	if len(matches) < 3 {
		return ""
	}
	return matches[2]
}

func (p Parser) Parse() (int, string) {

	id := p.parseId()
	subject := p.parseSubject()

	return id, subject
}

func (p Parser) Subject() string {
	if strings.Contains(p.input, " due") {
		index := strings.LastIndex(p.input, " due")
		return p.input[0:index]
	} else {
		return p.input
	}
}

func (p Parser) ExpandProject(input string) string {
	r := regexp.MustCompile(`(\+[\p{L}\d_-]+):`)
	matches := r.FindStringSubmatch(input)
	if len(matches) < 2 {
		return ""
	}

	return matches[1]
}

func (p Parser) Projects() []string {
	r := regexp.MustCompile(`\+[\p{L}\d_-]+`)
	return p.matchWords(p.input, r)
}

func (p Parser) Contexts() []string {
	r := regexp.MustCompile(`\@[\p{L}\d_]+`)
	return p.matchWords(p.input, r)
}

func (p Parser) hasDue() bool {
	r1 := regexp.MustCompile(`due \w+$`)
	r2 := regexp.MustCompile(`due \w+ \d+$`)
	return (r1.MatchString(p.input) || r2.MatchString(p.input))
}

func (p Parser) dueDate(pivot time.Time, input string) (time.Time, error) {
	date, err := weekdays.English(input, pivot).Weekday()
	if err != nil {
		log.Println(err)
		return date, err
	}

	return date, nil
}

// Due returns the parsed date. if any, formatted
// as a string. The given day is taken as a reference
// for relative dates (e.g. monday)
func (p Parser) Due(day time.Time) string {
	r := regexp.MustCompile(`due (.*)$`)
	matches := r.FindStringSubmatch(p.input)

	if len(matches) < 2 || matches[1] == "none" {
		return ""
	}

	if date, err := p.dueDate(day, matches[1]); err == nil {
		return date.Format("2006-01-02")
	}

	return p.parseArbitraryDate(matches[1], day)
}

func (p Parser) parseArbitraryDate(_date string, pivot time.Time) string {
	d1 := p.parseArbitraryDateWithYear(_date, pivot.Year())

	var diff1 time.Duration
	if d1.After(time.Now()) {
		diff1 = d1.Sub(pivot)
	} else {
		diff1 = pivot.Sub(d1)
	}
	d2 := p.parseArbitraryDateWithYear(_date, pivot.Year()+1)
	if d2.Sub(pivot) > diff1 {
		return d1.Format("2006-01-02")
	} else {
		return d2.Format("2006-01-02")
	}
}

func (p Parser) parseArbitraryDateWithYear(_date string, year int) time.Time {
	res := strings.Join([]string{_date, strconv.Itoa(year)}, " ")
	if date, err := time.Parse("Jan 2 2006", res); err == nil {
		return date
	}

	if date, err := time.Parse("2 Jan 2006", res); err == nil {
		return date
	}
	fmt.Printf("Could not parse the date you gave me: %s\n", _date)
	fmt.Println("I'm expecting a date like \"Dec 22\" or \"22 Dec\".")
	fmt.Println("See http://todolist.site/#adding for more info.")
	os.Exit(-1)
	return time.Now()
}

func (p Parser) matchWords(input string, r *regexp.Regexp) []string {
	results := r.FindAllString(input, -1)
	ret := []string{}

	for _, val := range results {
		ret = append(ret, val[1:])
	}
	return ret
}
