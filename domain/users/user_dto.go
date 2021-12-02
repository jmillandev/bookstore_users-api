package users

import (
	"strings"

	"github.com/jmillandev/bookstore_utils-go/rest_errors"
)

type User struct {
	Id          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
	Status      string `json:"status"`
	Password    string `json:"password"`
}

type Users []User

const (
	StatusActive = "active"
)

func (user *User) Validate() *rest_errors.RestErr {
	user.FirstName = strings.TrimSpace(user.FirstName)
	user.LastName = strings.TrimSpace(user.LastName)

	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == "" {
		return rest_errors.NewBadRequestError("Invalid email addres")
	}

	user.FirstName = strings.TrimSpace(strings.ToLower(user.FirstName))
	if user.Email == "" {
		return rest_errors.NewBadRequestError("Invalid first name")
	}

	user.LastName = strings.TrimSpace(strings.ToLower(user.LastName))
	if user.LastName == "" {
		return rest_errors.NewBadRequestError("Invalid last name")
	}

	user.Password = strings.TrimSpace(user.Password)
	if user.Password == "" {
		return rest_errors.NewBadRequestError("Invalid email password")
	}

	return nil
}
