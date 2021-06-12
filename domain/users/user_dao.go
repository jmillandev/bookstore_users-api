package users

import (
	"fmt"

	"github.com/jgmc3012/bookstore_users-api/datasources/mysql/users_db"
	"github.com/jgmc3012/bookstore_users-api/utils/date_utils"
	"github.com/jgmc3012/bookstore_users-api/utils/errors"
)

var (
	usersDB = make(map[int64]*User)
)

func Get(userId int64) (*User, *errors.RestErr) {
	if err := users_db.Client.Ping(); err != nil {
		panic(err)
	}
	user := usersDB[userId]
	if user == nil {
		return user, errors.NewNotFoundError(fmt.Sprintf("User %d not found", userId))
	}
	return user, nil
}

func (user *User) Save() *errors.RestErr {
	current_user := usersDB[user.Id]
	if current_user != nil {
		if current_user.Email == user.Email {
			return errors.NewBadRequestError(fmt.Sprintf("email %s already registered", user.Email))
		}
		return errors.NewBadRequestError(fmt.Sprintf("user %d already exists", user.Id))
	}

	user.DateCreated = date_utils.GetNowString()

	usersDB[user.Id] = user
	return nil
}
