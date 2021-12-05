package users

import (
	"strings"

	"github.com/jmillandev/bookstore_users-api/datasources/mysql/users_db"
	"github.com/jmillandev/bookstore_users-api/utils/mysql_utils"
	"github.com/jmillandev/bookstore_utils-go/logger"
	"github.com/jmillandev/bookstore_utils-go/rest_errors"
)

const (
	queryInsertUser             = "INSERT INTO users(first_name, last_name, email, status, password, date_created) VALUES(?, ?, ?, ?, ?, ?);"
	queryGetUser                = "SELECT id, first_name, last_name, email, status, date_created FROM users WHERE id=?"
	queryUpdateUser             = "UPDATE users SET first_name=?, last_name=?, email=? WHERE id=?"
	queryDeleteUser             = "DELETE FROM users WHERE id=?"
	queryFindUserByStatus       = "SELECT id, first_name, last_name, email, date_created, status FROM users WHERE status=?"
	queryFindByEmailAndPassword = "SELECT id, first_name, last_name, email, status, date_created FROM users WHERE email=? AND password=? AND status=?"
)

func (user *User) Get() *rest_errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryGetUser)
	if err != nil {
		logger.Error("error when trying to prepare get user statement", err)
		return rest_errors.NewInternalServerError("database error", err)
	}
	defer stmt.Close()

	result := stmt.QueryRow(user.Id)

	// results, err := stmt.Query(userId)

	if err := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.Status, &user.DateCreated); err != nil {
		logger.Error("error when trying to get user by id", err)
		return mysql_utils.ParseError(err)
	}

	return nil
}

func (user *User) Save() *rest_errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryInsertUser)
	if err != nil {
		logger.Error("error when trying to prepare save user statement", err)
		return rest_errors.NewInternalServerError("database error", err)
	}
	defer stmt.Close()

	insertResult, err := stmt.Exec(user.FirstName, user.LastName, user.Email, user.Status, user.Password, user.DateCreated)
	if err != nil {
		logger.Error("error when trying to save user", err)
		return mysql_utils.ParseError(err)
	}

	// insertResult, err := users_db.Client.Exec(queryInsertUser, user.FirstName, user.LastName, user.Email, user.DateCreated) // lost performance

	userId, err := insertResult.LastInsertId()
	if err != nil {
		logger.Error("error when trying to get last insert id after creating a new user", err)
		return mysql_utils.ParseError(err)
	}

	user.Id = userId
	return nil
}

func (user *User) Update() *rest_errors.RestErr {

	stmt, err := users_db.Client.Prepare(queryUpdateUser)
	if err != nil {
		logger.Error("error when trying to prepare update user statement", err)
		return rest_errors.NewInternalServerError("database error", err)
	}
	defer stmt.Close()

	if _, err = stmt.Exec(user.FirstName, user.LastName, user.Email, user.Id); err != nil {
		logger.Error("error when trying to update user", err)
		return mysql_utils.ParseError(err)
	}
	return nil

}

func (user *User) Delete() *rest_errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryDeleteUser)
	if err != nil {
		logger.Error("error when trying to prepare delete user statement", err)
		return rest_errors.NewInternalServerError("database error", err)
	}
	defer stmt.Close()

	if _, err = stmt.Exec(user.Id); err != nil {
		logger.Error("error when trying to delete user", err)
		return mysql_utils.ParseError(err)
	}
	return nil
}

func (user *User) FindByStatus(status string) ([]User, *rest_errors.RestErr) {
	stmt, err := users_db.Client.Prepare(queryFindUserByStatus)
	if err != nil {
		logger.Error("error when trying to prepare find user by status statement", err)
		return nil, rest_errors.NewInternalServerError("database error", err)
	}
	defer stmt.Close()

	rows, err := stmt.Query(status)
	if err != nil {
		logger.Error("error when trying to find user by status", err)
		return nil, mysql_utils.ParseError(err)
	}
	defer rows.Close()

	results := make([]User, 0)
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); err != nil {
			logger.Error("error when scan user row into user struct", err)
			return nil, mysql_utils.ParseError(err)
		}
		results = append(results, user)
	}
	return results, nil
}

func (user *User) FindByEmailAndPassword() *rest_errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryFindByEmailAndPassword)
	if err != nil {
		logger.Error("error when trying to prepare get user by email and password statement", err)
		return rest_errors.NewInternalServerError("database error", err)
	}
	defer stmt.Close()

	result := stmt.QueryRow(user.Email, user.Password, StatusActive)

	if err := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.Status, &user.DateCreated); err != nil {
		if strings.Contains(err.Error(), mysql_utils.ErrorNoRows) {
			return rest_errors.NewBadRequestError("invalid user credentials.")
		}
		logger.Error("error when trying to get user by email and password", err)
		return mysql_utils.ParseError(err)
	}

	return nil
}
