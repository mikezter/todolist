package main

import (
	"os"
	"strings"
	"testing"

	"github.com/gammons/todolist/todolist"
)

const testFile = ".todos.json"

func newTestApp() *todolist.App {
	removeTestFile()

	store := todolist.FileStore{FileLocation: testFile}
	store.Initialize()

	return &todolist.App{
		TodoList:  &todolist.TodoList{},
		TodoStore: &store,
	}
}

func removeTestFile() {
	os.Remove(testFile)
}

func TestExpandTodo(t *testing.T) {
	app := newTestApp()
	defer removeTestFile()

	err := routeInput(app, "a", "foobar")
	if err != nil {
		t.Fatal(err)
	}

	subTasks := []string{"get a website", "convert leads", "profit"}
	newProject := "+newProject"

	err = routeInput(app, "ex", "1 "+newProject+": "+strings.Join(subTasks, ", "))
	if err != nil {
		t.Fatal(err)
	}

	expected := [3]string{}

	for i, t := range subTasks {
		expected[i] = newProject + " " + t
	}

	for i, todo := range app.TodoList.Todos() {
		if todo.Subject != expected[i] {
			t.Fatal(todo.Subject, "!=", expected[i])
		}

	}

}

func TestAddTodo(t *testing.T) {
	app := newTestApp()
	defer removeTestFile()

	err := routeInput(app, "a", "foobar")
	if err != nil {
		t.Fatal(err)
	}

	for _, todo := range app.TodoList.Todos() {
		if todo.Subject == "foobar" {
			return
		}
	}

	t.Fatal("Todo not found")
}

func TestEditSubject(t *testing.T) {
	app := newTestApp()
	defer removeTestFile()

	err := routeInput(app, "a", "a foobar")
	if err != nil {
		t.Fatal(err)
	}

	err = routeInput(app, "es", "es 1 foobilicious")
	if err != nil {
		t.Fatal(err)
	}

	for _, todo := range app.TodoList.Todos() {
		if todo.Subject == "foobilicious" {
			return
		}
	}

	t.Fatal("Todo not found")
}
