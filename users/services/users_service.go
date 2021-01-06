package services

import (
	"smartshader/go-bookstore/users/domain/users"
	"smartshader/go-bookstore/users/utils/errors"
)

func GetUser(userId int64) (*users.User, *errors.RestErr) {
	res := &users.User{Id: userId}
	if err := res.Get(); err != nil {
		return nil, err
	}

	return res, nil
}

func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	if err := user.Save(); err != nil {
		return nil, err
	}

	return &user, nil
}
