package router

// This is a generated file
// Manual changes will be overwritten

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/julienschmidt/httprouter"
	log "github.com/sirupsen/logrus"

	"github.com/vincentvanderweele/swagger-todo-mvc/generated/model"
)

// Handler implements the actual functionality of the service
type Handler interface {
	// Other
	GetTodos() (model.ReadOnlyTodos, error)
	CreateTodo(bodyTodo model.Todo) (model.ReadOnlyTodo, error)
	GetTodo(pathTodoID string) (model.ReadOnlyTodo, error)
	UpdateTodo(pathTodoID string, bodyUpdate model.Todo) error
	DeleteTodo(pathTodoID string) error
	SetDone(pathTodoID string) error
}

// ErrorTransformer transforms an error into a message and code that can be returned via http
type ErrorTransformer interface {
	Transform(err error) (message string, code int)
}

type middleware struct {
	handler          Handler
	errorTransformer ErrorTransformer
}

// NewServer creates a http handler with a router for all methods of the service
func NewServer(handler Handler, errorTransformer ErrorTransformer) http.Handler {
	m := &middleware{
		handler,
		errorTransformer,
	}

	router := httprouter.New()

	router.GET("/todos", m.getTodos)
	router.POST("/todos", m.createTodo)
	router.GET("/todos/:todoId", m.getTodo)
	router.PUT("/todos/:todoId", m.updateTodo)
	router.DELETE("/todos/:todoId", m.deleteTodo)
	router.PUT("/todos/:todoId/setdone", m.setDone)

	return Recoverer(router)
}

// Recoverer handles unexpected panics and returns internal server error
func Recoverer(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.WithField("error", err).Error("Recovered")
				http.Error(w, "Internal server error", http.StatusInternalServerError)
			}
		}()
		next.ServeHTTP(w, r)
	})
}

func (m *middleware) getTodos(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var (
		result model.ReadOnlyTodos
		err    error
		errors []string
	)

	if len(errors) > 0 {
		log.WithFields(log.Fields{
			"handler": "getTodos",
			"errors":  strings.Join(errors, "\n"),
		})
		http.Error(w, strings.Join(errors, "\n"), http.StatusBadRequest)
		return
	}

	if result, err = m.handler.GetTodos(); err != nil {
		message, code := m.errorTransformer.Transform(err)
		http.Error(w, message, code)
		return
	}

	if errors = result.Validate(); len(errors) > 0 {
		log.WithFields(log.Fields{
			"dataType": "ReadOnlyTodos",
			"error":    strings.Join(errors, "\n"),
		}).Error("Invalid response data")
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	respondJSON(w, result, "ReadOnlyTodos")
}

func (m *middleware) createTodo(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var (
		result model.ReadOnlyTodo
		err    error
		errors []string
	)

	var bodyTodo model.Todo
	if err = json.NewDecoder(r.Body).Decode(&bodyTodo); err != nil {
		errors = append(errors, err.Error())
		log.WithFields(log.Fields{
			"bodyType": "Todo",
			"error":    err,
		}).Error("Failed to parse body data")
	} else if e := bodyTodo.Validate(); len(e) > 0 {
		errors = append(errors, e...)
	}
	if len(errors) > 0 {
		log.WithFields(log.Fields{
			"handler": "createTodo",
			"errors":  strings.Join(errors, "\n"),
		})
		http.Error(w, strings.Join(errors, "\n"), http.StatusBadRequest)
		return
	}

	if result, err = m.handler.CreateTodo(bodyTodo); err != nil {
		message, code := m.errorTransformer.Transform(err)
		http.Error(w, message, code)
		return
	}

	if errors = result.Validate(); len(errors) > 0 {
		log.WithFields(log.Fields{
			"dataType": "ReadOnlyTodo",
			"error":    strings.Join(errors, "\n"),
		}).Error("Invalid response data")
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	respondJSON(w, result, "ReadOnlyTodo")
}

func (m *middleware) getTodo(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	var (
		result model.ReadOnlyTodo
		err    error
		errors []string
	)

	pathTodoID := params.ByName("todoId")

	if len(errors) > 0 {
		log.WithFields(log.Fields{
			"handler": "getTodo",
			"errors":  strings.Join(errors, "\n"),
		})
		http.Error(w, strings.Join(errors, "\n"), http.StatusBadRequest)
		return
	}

	if result, err = m.handler.GetTodo(pathTodoID); err != nil {
		message, code := m.errorTransformer.Transform(err)
		http.Error(w, message, code)
		return
	}

	if errors = result.Validate(); len(errors) > 0 {
		log.WithFields(log.Fields{
			"dataType": "ReadOnlyTodo",
			"error":    strings.Join(errors, "\n"),
		}).Error("Invalid response data")
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	respondJSON(w, result, "ReadOnlyTodo")
}

func (m *middleware) updateTodo(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	var (
		err    error
		errors []string
	)

	pathTodoID := params.ByName("todoId")

	var bodyUpdate model.Todo
	if err = json.NewDecoder(r.Body).Decode(&bodyUpdate); err != nil {
		errors = append(errors, err.Error())
		log.WithFields(log.Fields{
			"bodyType": "Todo",
			"error":    err,
		}).Error("Failed to parse body data")
	} else if e := bodyUpdate.Validate(); len(e) > 0 {
		errors = append(errors, e...)
	}
	if len(errors) > 0 {
		log.WithFields(log.Fields{
			"handler": "updateTodo",
			"errors":  strings.Join(errors, "\n"),
		})
		http.Error(w, strings.Join(errors, "\n"), http.StatusBadRequest)
		return
	}

	if err = m.handler.UpdateTodo(pathTodoID, bodyUpdate); err != nil {
		message, code := m.errorTransformer.Transform(err)
		http.Error(w, message, code)
		return
	}

	w.Write([]byte("OK"))
}

func (m *middleware) deleteTodo(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	var (
		err    error
		errors []string
	)

	pathTodoID := params.ByName("todoId")

	if len(errors) > 0 {
		log.WithFields(log.Fields{
			"handler": "deleteTodo",
			"errors":  strings.Join(errors, "\n"),
		})
		http.Error(w, strings.Join(errors, "\n"), http.StatusBadRequest)
		return
	}

	if err = m.handler.DeleteTodo(pathTodoID); err != nil {
		message, code := m.errorTransformer.Transform(err)
		http.Error(w, message, code)
		return
	}

	w.Write([]byte("OK"))
}

func (m *middleware) setDone(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	var (
		err    error
		errors []string
	)

	pathTodoID := params.ByName("todoId")

	if len(errors) > 0 {
		log.WithFields(log.Fields{
			"handler": "setDone",
			"errors":  strings.Join(errors, "\n"),
		})
		http.Error(w, strings.Join(errors, "\n"), http.StatusBadRequest)
		return
	}

	if err = m.handler.SetDone(pathTodoID); err != nil {
		message, code := m.errorTransformer.Transform(err)
		http.Error(w, message, code)
		return
	}

	w.Write([]byte("OK"))
}

func respondJSON(w http.ResponseWriter, data interface{}, dataType string) {
	json, err := json.Marshal(data)
	if err != nil {
		log.WithFields(log.Fields{
			"dataType": dataType,
			"error":    err.Error(),
		}).Error("Failed to convert to json")
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}

func validateString(s, name string, minLength, maxLength *int, enum []string) (errors []string) {
	if minLength != nil {
		if len(s) < *minLength {
			errors = append(errors, fmt.Sprintf("%s should be no shorter than %d characters", name, *minLength))
		}
	}

	if maxLength != nil {
		if len(s) > *maxLength {
			errors = append(errors, fmt.Sprintf("%s should be no longer than %d characters", name, *maxLength))
		}
	}

	if enum != nil {
		found := false
		for i := range enum {
			if s == enum[i] {
				found = true
				break
			}
		}
		if !found {
			errors = append(errors, fmt.Sprintf("%s is not an allowed value for %s", s, name))
		}
	}

	return
}

func validateArray(a []string, name string, minItems, maxItems *int, uniqueItems bool) (errors []string) {
	if minItems != nil {
		if len(a) < *minItems {
			errors = append(errors, fmt.Sprintf("%s should have no less than %d elements", name, *minItems))
		}
	}

	if maxItems != nil {
		if len(a) > *maxItems {
			errors = append(errors, fmt.Sprintf("%s should have no more than %d elements", name, *maxItems))
		}
	}

	if uniqueItems {
		seen := map[string]struct{}{}
		for _, elt := range a {
			if _, duplicate := seen[elt]; duplicate {
				errors = append(errors, fmt.Sprintf("%s occurs multiple times in %s", elt, name))
			}
			seen[elt] = struct{}{}
		}
	}

	return
}
