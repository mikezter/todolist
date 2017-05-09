package todolist

import (
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestParseSubject(t *testing.T) {
	input := "do this thing"
	parser := Parser{input}
	todo := parser.ParseNewTodo()
	if todo.Subject != "do this thing" {
		t.Error("Expected todo.Subject to equal 'do this thing'")
	}
}

func TestParseSubjectWithDue(t *testing.T) {
	input := "do this thing due tomorrow"
	parser := Parser{input}
	todo := parser.ParseNewTodo()
	if todo.Subject != "do this thing" {
		t.Error("Expected todo.Subject to equal 'do this thing', got ", todo.Subject)
	}
}

func TestParseExpandProjects(t *testing.T) {
	assert := assert.New(t)
	parser := Parser{}
	correctFormat := parser.ExpandProject("ex 113 +meeting: figures, slides, coffee, suger")
	assert.Equal("+meeting", correctFormat)
	wrongFormat1 := parser.ExpandProject("ex 114 +meeting figures, slides, coffee, suger")
	assert.Equal("", wrongFormat1)
	wrongFormat2 := parser.ExpandProject("ex 115 meeting: figures, slides, coffee, suger")
	assert.Equal("", wrongFormat2)
	wrongFormat3 := parser.ExpandProject("ex 116 meeting figures, slides, coffee, suger")
	assert.Equal("", wrongFormat3)
	wrongFormat4 := parser.ExpandProject("ex 117 +重要な會議: 図, コーヒー, 砂糖")
	assert.Equal("+重要な會議", wrongFormat4)
}

func TestParseProjects(t *testing.T) {
	input := "do this thing +proj1 +proj2 +專案3 +proj-name due tomorrow"
	parser := Parser{input}
	todo := parser.ParseNewTodo()
	if len(todo.Projects) != 4 {
		t.Error("Expected Projects length to be 3")
	}
	if todo.Projects[0] != "proj1" {
		t.Error("todo.Projects[0] should equal 'proj1' but got", todo.Projects[0])
	}
	if todo.Projects[1] != "proj2" {
		t.Error("todo.Projects[1] should equal 'proj2' but got", todo.Projects[1])
	}
	if todo.Projects[2] != "專案3" {
		t.Error("todo.Projects[2] should equal '專案3' but got", todo.Projects[2])
	}
	if todo.Projects[3] != "proj-name" {
		t.Error("todo.Projects[3] should equal 'proj-name' but got", todo.Projects[3])
	}
}

func TestParseContexts(t *testing.T) {
	input := "do this thing with @bob and @mary due tomorrow"
	parser := Parser{input}
	todo := parser.ParseNewTodo()
	if len(todo.Contexts) != 2 {
		t.Error("Expected Projects length to be 2")
	}
	if todo.Contexts[0] != "bob" {
		t.Error("todo.Contexts[0] should equal 'mary' but got", todo.Contexts[0])
	}
	if todo.Contexts[1] != "mary" {
		t.Error("todo.Contexts[1] should equal 'mary' but got", todo.Contexts[1])
	}
}

func TestDueToday(t *testing.T) {
	parser := Parser{"do this thing with @bob and @mary due today"}
	todo := parser.ParseNewTodo()
	if todo.Due != bod(time.Now()).Format("2006-01-02") {
		fmt.Println("Date is different", todo.Due, time.Now())
	}

	parser.input = "do this thing with @bob and @mary due tod"
	todo = parser.ParseNewTodo()
	if todo.Due != bod(time.Now()).Format("2006-01-02") {
		fmt.Println("Date is different", todo.Due, time.Now())
	}
}

func TestDueTomorrow(t *testing.T) {
	parser := Parser{"do this thing with @bob and @mary due tomorrow"}
	todo := parser.ParseNewTodo()
	if todo.Due != bod(time.Now()).AddDate(0, 0, 1).Format("2006-01-02") {
		fmt.Println("Date is different", todo.Due, time.Now())
	}

	parser.input = "do this thing with @bob and @mary due tom"
	todo = parser.ParseNewTodo()
	if todo.Due != bod(time.Now()).AddDate(0, 0, 1).Format("2006-01-02") {
		fmt.Println("Date is different", todo.Due, time.Now())
	}
}

func TestDueSpecific(t *testing.T) {
	assert := assert.New(t)
	parser := Parser{"do this thing with @bob and @mary due jun 1"}
	todo := parser.ParseNewTodo()
	year := strconv.Itoa(time.Now().Year())
	assert.Equal(fmt.Sprintf("%s-06-01", year), todo.Due)
}

func TestMondayOnSunday(t *testing.T) {
	assert := assert.New(t)
	parser := Parser{}
	now, _ := time.Parse("2006-01-02", "2016-04-24")
	assert.Equal("2016-04-25", parser.monday(now))
}

func TestMondayOnMonday(t *testing.T) {
	assert := assert.New(t)
	parser := Parser{}
	now, _ := time.Parse("2006-01-02", "2016-04-25")
	assert.Equal("2016-04-25", parser.monday(now))
}

func TestMondayOnTuesday(t *testing.T) {
	assert := assert.New(t)
	parser := Parser{}
	now, _ := time.Parse("2006-01-02", "2016-04-26")
	assert.Equal("2016-05-02", parser.monday(now))
}

func TestTuesdayOnMonday(t *testing.T) {
	assert := assert.New(t)
	parser := Parser{}
	now, _ := time.Parse("2006-01-02", "2016-04-25")
	assert.Equal("2016-04-26", parser.tuesday(now))
}

func TestTuesdayOnWednesday(t *testing.T) {
	assert := assert.New(t)
	parser := Parser{}
	now, _ := time.Parse("2006-01-02", "2016-04-27")
	assert.Equal("2016-05-03", parser.tuesday(now))
}

func TestDueOnSpecificDate(t *testing.T) {
	assert := assert.New(t)
	year := time.Now().Year()

	parser := Parser{"due may 2"}
	assert.Equal(fmt.Sprintf("%v-05-02", year), parser.Due(time.Now()))

	parser.input = "due jun 1"
	assert.Equal(fmt.Sprintf("%v-06-01", year), parser.Due(time.Now()))
}

func TestDueOnSpecificDateEuropean(t *testing.T) {
	assert := assert.New(t)
	parser := Parser{"due 2 may"}
	year := time.Now().Year()
	assert.Equal(fmt.Sprintf("%v-05-02", year), parser.Due(time.Now()))
}

func TestDueIntelligentlyChoosesCorrectYear(t *testing.T) {
	assert := assert.New(t)
	parser := Parser{}
	marchTime, _ := time.Parse("2006-01-02", "2016-03-25")
	januaryTime, _ := time.Parse("2006-01-02", "2016-01-05")
	septemberTime, _ := time.Parse("2006-01-02", "2016-09-25")
	decemberTime, _ := time.Parse("2006-01-02", "2016-12-25")

	assert.Equal("2016-01-10", parser.parseArbitraryDate("jan 10", januaryTime))
	assert.Equal("2016-01-10", parser.parseArbitraryDate("jan 10", marchTime))
	assert.Equal("2017-01-10", parser.parseArbitraryDate("jan 10", septemberTime))
	assert.Equal("2017-01-10", parser.parseArbitraryDate("jan 10", decemberTime))
}

func TestParseCommandIdSubject(t *testing.T) {
	assert := assert.New(t)
	parser := Parser{"es 24 a new subject"}
	id, subject := parser.Parse()

	assert.Equal(24, id)
	assert.Equal("a new subject", subject)
}

func TestParseInvalidCommandIdSubject(t *testing.T) {
	assert := assert.New(t)
	input := "es a new project"
	parser := Parser{input}
	id, subject := parser.Parse()

	assert.Equal(-1, id)
	assert.Equal("", subject)
}

func newDate(y, m, d int) time.Time {
	return time.Date(y, time.Month(m), d, 0, 0, 0, 1, time.Local)
}
func TestDueWeekdays(t *testing.T) {

	today := newDate(2017, 5, 9) // tuesday
	dates := map[string]time.Time{
		"due tod":       newDate(2017, 5, 9),  // today
		"due tom":       newDate(2017, 5, 10), // tomorrow
		"due tue":       newDate(2017, 5, 9),  // tuesday
		"due wed":       newDate(2017, 5, 10), // wednesday
		"due thu":       newDate(2017, 5, 11), // thursday
		"due fri":       newDate(2017, 5, 12), // friday
		"due sat":       newDate(2017, 5, 13), // saturday
		"due sun":       newDate(2017, 5, 14), // sunday
		"due mon":       newDate(2017, 5, 15), // monday
		"due last week": newDate(2017, 5, 1),  // monday last week
		"due next week": newDate(2017, 5, 15), // monday next week
	}

	p := Parser{}

	for input, date := range dates {
		p.input = input

		expected := date.Format("2006-01-02")
		actual := p.Due(today)

		if actual != expected {
			t.Error(input, actual, expected)
		}
	}
}

func TestDueInvalid(t *testing.T) {
	today := newDate(2017, 5, 9) // tuesday
	p := Parser{"no asd"}
	if p.Due(today) != "" {
		t.Fail()
	}
}

func TestDueNone(t *testing.T) {
	today := newDate(2017, 5, 9) // tuesday
	p := Parser{"due none"}
	if p.Due(today) != "" {
		t.Fail()
	}

}
