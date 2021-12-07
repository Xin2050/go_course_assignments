package users

import (
	mysql_db "github.com/Xin2050/go_course_assignments/s1/datasources/mysql"
	"github.com/Xin2050/go_course_assignments/s1/logger"
	"github.com/Xin2050/go_course_assignments/s1/utils/errors"
	mysql_utils "github.com/Xin2050/go_course_assignments/s1/utils/mysql"
)

const (
	queryGetUser = "Select id, lastName from users where id=?;"
)

func (user *User) Get() *errors.RestError {
	stmt, err := mysql_db.Client.Prepare(queryGetUser)
	if err != nil {
		logger.Error("error when trying to prepare get user statement", err)
		return mysql_utils.ParseError(err)
	}
	defer stmt.Close()
	err = stmt.QueryRow(user.Id).Scan(&user.Id, &user.Name)
	if err != nil {
		logger.Error("error when trying to get user by id", err)
		return mysql_utils.ParseError(err)
	}
	return nil
}
