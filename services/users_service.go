package services

import (
	"github.com/jmillandev/bookstore_users-api/domain/users"
	"github.com/jmillandev/bookstore_users-api/utils/date_utils"
	"github.com/jmillandev/bookstore_utils-go/cripto_utils"
	"github.com/jmillandev/bookstore_utils-go/rest_errors"
)

type usersService struct{}

type usersServiceInterface interface {
	GetUser(users.User) (*users.User, *rest_errors.RestErr)
	CreateUser(users.User) (*users.User, *rest_errors.RestErr)
	UpdateUser(bool, users.User) (*users.User, *rest_errors.RestErr)
	DeleteUser(int64) *rest_errors.RestErr
	FindByStatusUser(string) (users.Users, *rest_errors.RestErr)
	LoginUser(users.LoginRequest) (*users.User, *rest_errors.RestErr)
}

var UsersService usersServiceInterface = &usersService{}

func (s *usersService) GetUser(user users.User) (*users.User, *rest_errors.RestErr) {
	currentUser := &users.User{Id: user.Id}

	if errUser := currentUser.Get(); errUser != nil {
		return nil, errUser
	}
	return currentUser, nil
}

func (s *usersService) CreateUser(user users.User) (*users.User, *rest_errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	user.DateCreated = date_utils.GetNowString()
	user.Status = users.StatusActive
	user.Password = cripto_utils.GetMd5(user.Password)
	if err := user.Save(); err != nil {
		return nil, err
	}

	return &user, nil
}

func (s *usersService) UpdateUser(isPartial bool, user users.User) (*users.User, *rest_errors.RestErr) {
	currentUser, err := UsersService.GetUser(user)
	if err != nil {
		return nil, err
	}

	if err := user.Validate(); err != nil {
		return nil, err
	}

	if isPartial {
		if user.FirstName != "" {
			currentUser.FirstName = user.FirstName
		}
		if user.LastName != "" {
			currentUser.LastName = user.LastName
		}
		if user.Email != "" {
			currentUser.Email = user.Email
		}
	} else {
		currentUser.FirstName = user.FirstName
		currentUser.LastName = user.LastName
		currentUser.Email = user.Email
	}

	if err := currentUser.Update(); err != nil {
		return nil, err
	}

	return currentUser, nil

}

func (s *usersService) DeleteUser(userId int64) *rest_errors.RestErr {
	user := &users.User{Id: userId}
	return user.Delete()
}

func (s *usersService) FindByStatusUser(status string) (users.Users, *rest_errors.RestErr) {
	user := &users.User{}
	return user.FindByStatus(status)
}

func (s *usersService) LoginUser(request users.LoginRequest) (*users.User, *rest_errors.RestErr) {
	user := &users.User{
		Email:    request.Email,
		Password: cripto_utils.GetMd5(request.Password),
	}

	if err := user.FindByEmailAndPassword(); err != nil {
		return nil, err
	}

	return user, nil
}
