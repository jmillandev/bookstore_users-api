package services

import (
	"github.com/jgmc3012/bookstore_users-api/domain/users"
	"github.com/jgmc3012/bookstore_users-api/utils/errors"
)

func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	return &user, nil
}
