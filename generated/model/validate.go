package model

// This is a generated file
// Manual changes will be overwritten

// Validate validates a Todo based on the swagger spec
func (s *Todo) Validate() (errors []string) {

	if s.Title == nil {
		errors = append(errors, "title is required")
	}

	if len(*s.Title) < 1 {
		errors = append(errors, "title should be no shorter than 1 characters")
	}
	return
}

// Validate validates a ReadOnlyTodo based on the swagger spec
func (s *ReadOnlyTodo) Validate() (errors []string) {
	if e := s.Todo.Validate(); len(e) > 0 {
		errors = append(errors, e...)
	}

	if s.Done == nil {
		errors = append(errors, "done is required")
	}

	if s.ID == nil {
		errors = append(errors, "id is required")
	}
	return
}

// Validate validates a Todos based on the swagger spec
func (s *Todos) Validate() (errors []string) {

	for _, elt := range *s {
		if e := elt.Validate(); len(e) > 0 {
			errors = append(errors, e...)
		}
	}
	return
}

// Validate validates a ReadOnlyTodos based on the swagger spec
func (s *ReadOnlyTodos) Validate() (errors []string) {

	for _, elt := range *s {
		if e := elt.Validate(); len(e) > 0 {
			errors = append(errors, e...)
		}
	}
	return
}
