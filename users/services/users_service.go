package services

import (
	"smartshader/go-bookstore/users/domain/users"
	"smartshader/go-bookstore/users/utils/errors"
)

func GetUser() {}

func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	return &user, nil
}
