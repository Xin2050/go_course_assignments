package users_services

import (
	"github.com/Xin2050/go_course_assignments/s1/domain/users"
	"github.com/Xin2050/go_course_assignments/s1/utils/errors"
)

type usersService struct{}

var (
	UsersService usersServiceInterface = &usersService{}
)

type usersServiceInterface interface {
	GetUser(int64) (*users.User, *errors.RestError)
}

func (s *usersService) GetUser(userId int64) (*users.User, *errors.RestError) {
	result := &users.User{Id: userId}
	if err := result.Get(); err != nil {
		return nil, err
	}
	return result, nil
}
