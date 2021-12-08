package users

import (
	mysql_db "github.com/Xin2050/go_course_assignments/s1/datasources/mysql"
	"github.com/pkg/errors"
)

const (
	queryGetUser    = "Select id, lastName from users where id=?;"
	queryInsertUser = "Insert into users(firstName, lastName, email) values (?,?,?);"
)

func (user *User) Save() error {
	stmt, err := mysql_db.Client.Prepare(queryInsertUser)
	if err != nil {
		return errors.Wrap(err, "error when trying to prepare save user statement")
	}
	defer stmt.Close()
	insertRs, saveErr := stmt.Exec(user.FirstName, user.LastName, user.Email)
	if saveErr != nil {
		return errors.Wrap(saveErr, "error when trying to save user statement")
	}
	userId, err := insertRs.LastInsertId()
	if err != nil {
		return errors.Wrap(err, "error when trying to get last insert user's id")
	}
	user.Id = userId
	return nil
}

func (user *User) Get() error {

	stmt, err := mysql_db.Client.Prepare(queryGetUser)
	if err != nil {
		//logger.Error("error when trying to prepare get user statement", err)
		return errors.Wrap(err, "User.Get() error when trying to prepare get user statement")
	}
	defer stmt.Close()
	queryErr := stmt.QueryRow(user.Id).Scan(&user.Id, &user.FirstName, &user.LastName)
	if queryErr != nil {
		//logger.Error("error when trying to get user by id", err)
		return errors.Wrap(queryErr, "User.Get() error when trying to get user by id")
	}
	return nil
}
