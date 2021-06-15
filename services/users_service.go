package services

import (
	"github.com/jgmc3012/bookstore_users-api/domain/users"
	"github.com/jgmc3012/bookstore_users-api/utils/errors"
)

func GetUser(user users.User) (*users.User, *errors.RestErr) {
	currentUser := &users.User{Id: user.Id}

	if errUser := currentUser.Get(); errUser != nil {
		return nil, errUser
	}
	return currentUser, nil
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

func UpdateUser(user users.User) (*users.User, *errors.RestErr) {
	currentUser, err := GetUser(user)
	if err != nil {
		return nil, err
	}

	if err := user.Validate(); err != nil {
		return nil, err
	}

	currentUser.FirstName = user.FirstName
	currentUser.LastName = user.LastName
	currentUser.Email = user.Email

	if err := currentUser.Update(); err != nil {
		return nil, err
	}

	return currentUser, nil

}
