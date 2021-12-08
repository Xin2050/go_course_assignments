package mysql_utils

import (
	"fmt"
	rest_errors "github.com/Xin2050/go_course_assignments/s1/utils/errors"
	"github.com/go-sql-driver/mysql"
)

func ParseError(err error) *rest_errors.RestError {

	sqlErr, ok := err.(*mysql.MySQLError)
	if !ok {
		return rest_errors.InternalServerError(err.Error())
	}
	switch sqlErr.Number {
	case 1062:
		return rest_errors.BadRequestError(fmt.Sprintf("duplicated key: %s", sqlErr.Message))
	}
	return rest_errors.InternalServerError(fmt.Sprintf("error processing request:%s", err))
}
