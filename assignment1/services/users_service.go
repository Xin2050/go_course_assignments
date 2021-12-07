package users_services

import (
	"github.com/Xin2050/go_course_assignments/s1/domain/users"
	"github.com/pkg/errors"
)

type usersService struct{}

var (
	UsersService usersServiceInterface = &usersService{}
)

type usersServiceInterface interface {
	GetUser(int64) (*users.User, error)
}

func (s *usersService) GetUser(userId int64) (*users.User, error) {
	result := &users.User{Id: userId}
	if err := result.Get(); err != nil {
		return nil, errors.Wrap(err, "users_services:GetUser error")
	}
	return result, nil
}
