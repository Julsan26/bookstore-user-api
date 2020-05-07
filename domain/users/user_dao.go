package users

import (
	"fmt"
	"github.com/julsan26/bookstore-user-api/controllers/datasources/mysql/user_db"
	"github.com/julsan26/bookstore-user-api/utils/errors"
)


var (
	userDB= make(map[int64]*User)
)

func (user *User)Get()(*errors.RestErr){

	if err:= user_db.Client.Ping();err!=nil{
		panic (err)
	}

	result:=userDB[user.Id]
	if result==nil{
		  errors.NewNotFound(fmt.Sprintf("user %d not foung",user.Id))
	}
	user.Id=result.Id
	user.Email=result.Email
	user.DateCreated=result.DateCreated
	user.FirstName=result.FirstName
	user.LastName=result.LastName
 	return nil
}

func (user *User)Save()*errors.RestErr{
	current:=userDB[user.Id]
	if current!=nil{
		if current.Email==user.Email{
			return  errors.NewBadRequestError(fmt.Sprintf("email %v already exist",user.Email))
		}
		return  errors.NewBadRequestError(fmt.Sprintf("user %d already exist",user.Id))
	}
	userDB[user.Id]=user
	return nil
}
