package mysql_utils

import (
	"database/sql"
	"fmt"
	"github.com/Xin2050/go_course_assignments/s1/utils/errors"
	"github.com/go-sql-driver/mysql"
)

func ParseError(err error) *errors.RestError {
	if err == sql.ErrNoRows {
		return errors.NotFoundError(fmt.Sprintf("no record found by given condictions"))
	}
	sqlErr, ok := err.(*mysql.MySQLError)
	if !ok {
		return errors.InternalServerError(fmt.Sprintf("mysql error: %s", err))
	}
	switch sqlErr.Number {
	case 1062:
		return errors.BadRequestError(fmt.Sprintf("duplicated key: %s", sqlErr.Message))
	}
	return errors.InternalServerError(fmt.Sprintf("error processing request:%s", err))
}
