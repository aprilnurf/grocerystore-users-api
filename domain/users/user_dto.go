package users

import (
	"github.com/aprilnurf/grocerystore_users-api/utils/errors_utils"
	"strings"
)

const (
	StatusActive = true
)

type User struct {
	Id          int64  `json:"id"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	Email       string `json:"email"`
	CreatedDate string `json:"createdDate"`
	Password    string `json:"password"`
	Status      bool   `json:"status"`
}

//func Validate(user *User) *errors.RestError {
//	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
//	if user.Email == "" {
//		return errors.NewBadRequestError("Invalid Email Address")
//	}
//	return nil
//}

func (user *User) Validate() *errors_utils.RestError {
	user.FirstName = strings.TrimSpace(user.FirstName)
	user.LastName = strings.TrimSpace(user.LastName)
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == "" {
		return errors_utils.NewBadRequestError("Invalid Email Address")
	}
	user.Password = strings.TrimSpace(user.Password)
	if user.Password == ""{
		return errors_utils.NewBadRequestError("Invalid Password")
	}
	return nil
}
