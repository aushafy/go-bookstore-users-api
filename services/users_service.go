package services

import (
	"github.com/aushafy/go-bookstore-users-api/domain/users"
	"github.com/aushafy/go-bookstore-users-api/utils/errors"
)

func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	if err := user.Save(); err != nil {
		return nil, err
	}
	return &user, nil
}
