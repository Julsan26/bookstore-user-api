package services

import (
	"github.com/julsan26/bookstore-user-api/domain/users"
	"github.com/julsan26/bookstore-user-api/utils/errors"

)

func CreateUser(user users.User)(*users.User , *errors.RestErr) {

	if err := user.Validate(); err != nil {
		return nil, err
	}

	if err := user.Save(); err != nil {
		return nil, err
	}


	return &user,nil
}
