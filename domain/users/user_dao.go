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
	errorNoRows      = "no rows in result set"
	queryInsertUser  = "INSERT INTO users(first_name, last_name, email, date_created) VALUES(?, ?, ?, ?);"
	queryGetUser     = "SELECT id, first_name, last_name, email, date_created FROM users WHERE id=?"
)

func Get(userId int64) (*User, *errors.RestErr) {
	stmt, err := users_db.Client.Prepare(queryGetUser)
	if err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	result := stmt.QueryRow(userId)

	// results, err := stmt.Query(userId)

	user := &User{}

	if err := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated); err != nil {
		if strings.Contains(err.Error(), errorNoRows) {
			return nil, errors.NewNotFoundError(
				fmt.Sprintf("user(id: %d) not found", userId),
			)
		}
		return nil, errors.NewInternalServerError(
			fmt.Sprintf("error when trying to get user(id: %d): %s", userId, err.Error()),
		)
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
