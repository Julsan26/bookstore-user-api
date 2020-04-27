package services

import (
	"github.com/julsan26/bookstore-user-api/domain/users"
	"github.com/julsan26/bookstore-user-api/utils/errors"
	"net/http"
)

func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	return &user, nil

	return nil, &errors.RestErr{
		Status: http.StatusInternalServerError,
	}

}
