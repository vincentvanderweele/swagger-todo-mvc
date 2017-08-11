package model

// This is a generated file
// Manual changes will be overwritten

// Todo A task that has to be done
type Todo struct {
	Title *string `json:"title"`
}

// ReadOnlyTodo A task that has to be done
type ReadOnlyTodo struct {
	Todo
	Done *bool   `json:"done"`
	ID   *string `json:"id"`
}

// Todos A list of todos
type Todos []Todo

// ReadOnlyTodos A list of todos
type ReadOnlyTodos []ReadOnlyTodo
