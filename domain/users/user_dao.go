package users

import (
	"fmt"
	"github.com/aprilnurf/grocerystore_users-api/utils/errors"
)

var (
	usersDB = make(map[int64]*User)
)

func (user *User) Get() *errors.RestError {
	result := usersDB[user.Id]
	if result == nil {
		return errors.NewNotExistError(fmt.Sprintf("User %d Not Exist", user.Id))
	}
	user.Id = result.Id
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.Email = result.Email
	user.CreatedDate = result.CreatedDate
	return nil
}

func (user *User) Save() *errors.RestError {
	current := usersDB[user.Id]
	if current != nil {
		if current.Email == user.Email {
			return errors.NewBadRequestError(fmt.Sprintf("email %s already register", user.Email))
		}
		return errors.NewBadRequestError(fmt.Sprintf("User %d already exist", user.Id))
	}
	usersDB[user.Id] = user
	return nil
}
