package users

import (
	mysql_db "github.com/Xin2050/go_course_assignments/s1/datasources/mysql"
	"github.com/pkg/errors"
)

const (
	queryGetUser = "Select id, lastName from users where id=?;"
)

func (user *User) Get() error {

	stmt, err := mysql_db.Client.Prepare(queryGetUser)
	if err != nil {
		//logger.Error("error when trying to prepare get user statement", err)
		return errors.Wrap(err, "User.Get() error when trying to prepare get user statement")
	}
	defer stmt.Close()
	queryErr := stmt.QueryRow(user.Id).Scan(&user.Id, &user.Name)
	if queryErr != nil {
		//logger.Error("error when trying to get user by id", err)
		return errors.Wrap(queryErr, "User.Get() error when trying to get user by id")
	}
	return nil
}
