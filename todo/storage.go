package todo

import (
	"strconv"

	"github.com/vincentvanderweele/swagger-todo-mvc/generated/model"
	"github.com/vincentvanderweele/swagger-todo-mvc/generated/router"
)

type todoStorage struct {
	todos  []model.ReadOnlyTodo
	nextID int
}

// NewStorage creates a new todo storage
func NewStorage() router.Handler {
	return &todoStorage{
		todos:  []model.ReadOnlyTodo{},
		nextID: 0,
	}
}

func (s *todoStorage) GetTodos() (model.ReadOnlyTodos, error) {
	return model.ReadOnlyTodos(s.todos), nil
}

func (s *todoStorage) CreateTodo(newTodo model.Todo) (todo model.ReadOnlyTodo, err error) {
	todo = model.ReadOnlyTodo{
		Todo: newTodo,
		ID:   strp(strconv.Itoa(s.nextID)),
		Done: boolp(false),
	}

	s.todos = append(s.todos, todo)
	s.nextID++

	return
}

func (s *todoStorage) GetTodo(todoID string) (todo model.ReadOnlyTodo, err error) {
	var t *model.ReadOnlyTodo
	if t, _, err = s.getTodo(todoID); err != nil {
		return
	}

	todo = *t
	return
}

func (s *todoStorage) UpdateTodo(todoID string, update model.Todo) (err error) {
	var todo *model.ReadOnlyTodo
	if todo, _, err = s.getTodo(todoID); err != nil {
		return
	}

	todo.Todo = update

	return
}

func (s *todoStorage) DeleteTodo(todoID string) (err error) {
	var index int
	if _, index, err = s.getTodo(todoID); err != nil {
		return
	}

	s.todos[index] = s.todos[len(s.todos)-1]
	s.todos = s.todos[:len(s.todos)-1]

	return
}

func (s *todoStorage) SetDone(todoID string) (err error) {
	var todo *model.ReadOnlyTodo
	if todo, _, err = s.getTodo(todoID); err != nil {
		return
	}

	todo.Done = boolp(true)

	return
}

func (s *todoStorage) getTodo(todoID string) (todo *model.ReadOnlyTodo, index int, err error) {
	for index = range s.todos {
		if *s.todos[index].ID == todoID {
			todo = &s.todos[index]
			return
		}
	}

	err = todoNotFoundError("")
	return
}

func strp(s string) *string {
	return &s
}

func boolp(b bool) *bool {
	return &b
}
