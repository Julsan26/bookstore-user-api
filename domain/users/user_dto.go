package users

import (
	"github.com/julsan26/bookstore-user-api/utils/errors"
	"strings"
)

type User struct {
	Id          int64  `json:"id"`
	FirstName   string `json:"firstname"`
	LastName    string `json:"lastname"`
	DateCreated int    `json:"datecreated"`
	Email 		string 	`json:"email"`
}

func (user *User)Validate()*errors.RestErr{
	user.Email=strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email==""{
		return errors.NewBadRequestError("invalid email address")
	}
	return nil

}