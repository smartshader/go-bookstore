package services

import (
	"smartshader/go-bookstore/users/domain/users"
	"smartshader/go-bookstore/users/utils/crypto_utils"
	"smartshader/go-bookstore/users/utils/date_utils"
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

	user.DateCreated = date_utils.GetNowDBFormat()
	user.Status = users.StatusActive
	user.Password = crypto_utils.GetMd5(user.Password)

	if err := user.Save(); err != nil {
		return nil, err
	}

	return &user, nil
}

func UpdateUser(isPartial bool, user users.User) (*users.User, *errors.RestErr) {
	current, restErr := GetUser(user.Id)
	if restErr != nil {
		return nil, restErr
	}

	if isPartial {
		if user.FirstName != "" {
			current.FirstName = user.FirstName
		}

		if user.LastName != "" {
			current.LastName = user.LastName
		}

		if user.Email != "" {
			current.Email = user.Email
		}
	} else {
		current.FirstName = user.FirstName
		current.LastName = user.LastName
		current.Email = user.Email
	}

	if restErr = current.Update(); restErr != nil {
		return nil, restErr
	}

	return current, nil
}

func DeleteUser(userId int64) *errors.RestErr {
	user := &users.User{Id: userId}
	return user.Delete()
}

func SearchUsers(status string) (users.Users, *errors.RestErr) {
	dao := &users.User{}
	return dao.FindByStatus(status)
}
