package users

import (
	"github.com/aprilnurf/grocerystore_users-api/utils/errors"
	"strings"
)

type User struct {
	Id          int64  `json:"id"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	Email       string `json:"email"`
	CreatedDate string `json:"createdDate"`
}

//func Validate(user *User) *errors.RestError {
//	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
//	if user.Email == "" {
//		return errors.NewBadRequestError("Invalid Email Address")
//	}
//	return nil
//}

func (user *User) Validate() *errors.RestError {
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == "" {
		return errors.NewBadRequestError("Invalid Email Address")
	}
	return nil
}
