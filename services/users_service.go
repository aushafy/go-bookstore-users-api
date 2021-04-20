package services

import "github.com/aushafy/go-bookstore-users-api/domain/users"

func CreateUser(user users.User) (*users.User, error) {
	return &user, nil
}
