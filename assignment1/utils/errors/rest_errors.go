package rest_errors

import (
	"database/sql"
	"fmt"
	"github.com/Xin2050/go_course_assignments/s1/logger"
	"github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
	"net/http"
)

type RestError struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
	Error   string `json:"error"`
}

func ParseError(err error) *RestError {
	//write log and print stack trace
	logger.Errors(err)

	// check sql errors
	if errors.Is(err, sql.ErrNoRows) {
		return NotFoundError(fmt.Sprintf("no record found by given condictions"))
	}

	// check original err if a mysql error
	switch err := errors.Cause(err).(type) {
	case *mysql.MySQLError:
		// break down the mysql error code , send to client a user-friendly message
		return ParseMySQLError(errors.Cause(err))
	}

	return InternalServerError("Internal Server Error")
}

func ParseMySQLError(err error) *RestError {

	sqlErr, ok := err.(*mysql.MySQLError)
	if !ok {
		return InternalServerError(err.Error())
	}
	switch sqlErr.Number {
	case 1364:
		return InternalServerError(fmt.Sprintf("mysql Schema error"))
	case 1062:
		return BadRequestError(fmt.Sprintf("duplicated key: %s", sqlErr.Message))
	}
	return InternalServerError(fmt.Sprintf("error processing request:%s", err))
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
