package todolist

import "regexp"

type TodoFilter struct {
	Todos []*Todo
}

func NewFilter(todos []*Todo) *TodoFilter {
	return &TodoFilter{Todos: todos}
}

func (f *TodoFilter) Filter(input string) []*Todo {
	f.Todos = f.filterArchived(input)
	f.Todos = f.filterPrioritized(input)
	f.Todos = f.filterProjects(input)
	f.Todos = f.filterContexts(input)
	f.Todos = NewDateFilter(f.Todos).FilterDate(input)

	return f.Todos
}

func (t *TodoFilter) isFilteringByProjects(input string) bool {
	parser := Parser{input}
	return len(parser.Projects()) > 0
}

func (t *TodoFilter) isFilteringByContexts(input string) bool {
	parser := Parser{input}
	return len(parser.Contexts()) > 0
}

func (f *TodoFilter) filterArchived(input string) []*Todo {
	r, _ := regexp.Compile(`l archived$`)
	if r.MatchString(input) {
		return f.getArchived()
	} else {
		return f.getUnarchived()
	}
}

func (f *TodoFilter) filterPrioritized(input string) []*Todo {
	r, _ := regexp.Compile(`l p`)
	if r.MatchString(input) {
		return f.getPrioritized()
	} else {
		return f.Todos
	}
}

func (f *TodoFilter) filterProjects(input string) []*Todo {
	if !f.isFilteringByProjects(input) {
		return f.Todos
	}
	parser := Parser{input}
	projects := parser.Projects()
	var ret []*Todo

	for _, todo := range f.Todos {
		for _, todoProject := range todo.Projects {
			for _, project := range projects {
				if project == todoProject {
					ret = AddTodoIfNotThere(ret, todo)
				}
			}
		}
	}
	return ret
}

func (f *TodoFilter) filterContexts(input string) []*Todo {
	if !f.isFilteringByContexts(input) {
		return f.Todos
	}
	parser := Parser{input}
	contexts := parser.Contexts()
	var ret []*Todo

	for _, todo := range f.Todos {
		for _, todoContext := range todo.Contexts {
			for _, context := range contexts {
				if context == todoContext {
					ret = AddTodoIfNotThere(ret, todo)
				}
			}
		}
	}
	return ret
}

func (f *TodoFilter) getArchived() []*Todo {
	var ret []*Todo
	for _, todo := range f.Todos {
		if todo.Archived {
			ret = append(ret, todo)
		}
	}
	return ret
}

func (f *TodoFilter) getPrioritized() []*Todo {
	var ret []*Todo
	for _, todo := range f.Todos {
		if todo.IsPriority {
			ret = append(ret, todo)
		}
	}
	return ret
}

func (f *TodoFilter) getUnarchived() []*Todo {
	var ret []*Todo
	for _, todo := range f.Todos {
		if !todo.Archived {
			ret = append(ret, todo)
		}
	}
	return ret
}
