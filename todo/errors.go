package todo

import (
	"net/http"

	"github.com/vincentvanderweele/swagger-todo-mvc/generated/router"
)

type todoNotFoundError string

func (e todoNotFoundError) Error() string {
	return string(e)
}

type errorTransformer struct{}

// NewErrorTransformer returns a new error transformer
func NewErrorTransformer() router.ErrorTransformer {
	return &errorTransformer{}
}

func (h *errorTransformer) Transform(err error) (message string, code int) {
	switch err.(type) {
	case todoNotFoundError:
		message = "Todo not found"
		code = http.StatusNotFound
	default:
		message = "Internal server error"
		code = http.StatusInternalServerError
	}

	return
}
