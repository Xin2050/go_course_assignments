package rest_errors

import (
	"database/sql"
	"fmt"
	"github.com/pkg/errors"
	"net/http"
)

type RestError struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
	Error   string `json:"error"`
}

func ParseError(err error) *RestError {
	if errors.Is(err, sql.ErrNoRows) {
		return NotFoundError(fmt.Sprintf("no record found by given condictions"))
	}
	//todo unwrap error chain until find mysql error, then user mysql_utils.ParseError to handle mysql error code
	//todo more different errors handling
	return InternalServerError("Internal Server Error")
}

func BadRequestError(message string) *RestError {
	return &RestError{
		Status:  http.StatusBadRequest,
		Message: message,
		Error:   "bad_request",
	}
}
func NotFoundError(message string) *RestError {
	return &RestError{
		Status:  http.StatusNotFound,
		Message: message,
		Error:   "not_found_error",
	}
}

func InternalServerError(message string) *RestError {
	return &RestError{
		Status:  http.StatusInternalServerError,
		Message: message,
		Error:   "internal_server_error",
	}
}
