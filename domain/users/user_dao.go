package users

import (
	"fmt"
	"strings"

	"github.com/jgmc3012/bookstore_users-api/datasources/mysql/users_db"
	"github.com/jgmc3012/bookstore_users-api/utils/date_utils"
	"github.com/jgmc3012/bookstore_users-api/utils/errors"
)

const (
	ixdexUniqueEmail = "email_UNIQUE"
	queryInsertUser  = "INSERT INTO users(first_name, last_name, email, date_created) VALUES(?, ?, ?, ?);"
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
	stmt, err := users_db.Client.Prepare(queryInsertUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()
	user.DateCreated = date_utils.GetNowString()
	insertResult, err := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated)
	if err != nil {
		if strings.Contains(err.Error(), ixdexUniqueEmail) {
			return errors.NewBadRequestError(
				fmt.Sprintf("email '%s' already exists", user.Email),
			)
		}
		return errors.NewInternalServerError(
			fmt.Sprintf("error when trying to save user: %s", err.Error()),
		)
	}

	// insertResult, err := users_db.Client.Exec(queryInsertUser, user.FirstName, user.LastName, user.Email, user.DateCreated) // lost performance

	userId, err := insertResult.LastInsertId()
	if err != nil {
		return errors.NewInternalServerError(
			fmt.Sprintf("error when trying to save user: %s", err.Error()),
		)
	}

	user.Id = userId
	return nil
}
