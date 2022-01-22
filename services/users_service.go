package services

import (
	"github.com/aprilnurf/grocerystore_users-api/domain/users"
	"github.com/aprilnurf/grocerystore_users-api/utils/crypto_utils"
	"github.com/aprilnurf/grocerystore_users-api/utils/date_utils"
	"github.com/aprilnurf/grocerystore_users-api/utils/errors_utils"
)

func CreateUser(user users.User) (*users.User, *errors_utils.RestError) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	user.Password = crypto_utils.GetMd5(user.Password)
	user.CreatedDate = date_utils.GetNowDB()
	if err := user.Save(); err != nil {
		return nil, err
	}
	return &user, nil
}

func GetUser(userId int64) (*users.User, *errors_utils.RestError) {
	if userId <= 0 {
		return nil, errors_utils.NewBadRequestError("userId invalid")
	}
	result := &users.User{Id: userId}
	if err := result.Get(); err != nil {
		return nil, err
	}
	return result, nil
}

func UpdateUser(isPartial bool, user users.User) (*users.User, *errors_utils.RestError) {
	current, err := GetUser(user.Id)
	if err != nil {
		return nil, err
	}
	if err := user.Validate(); err != nil {
		return nil, err
	}
	if isPartial {
		if user.Email != "" {
			current.Email = user.Email
		}
		if user.FirstName != "" {
			current.FirstName = user.FirstName
		}
		if user.LastName != "" {
			current.LastName = user.LastName
		}
	} else {
		current.FirstName = user.FirstName
		current.LastName = user.LastName
		current.Email = user.Email
		current.Status = user.Status
	}
	if err := current.Update(); err != nil {
		return nil, err
	}
	return current, nil
}

func Search(status bool) ([]users.User, *errors_utils.RestError) {
	dao := &users.User{}
	return dao.FindByStatus(status)
}

func DeleteUser(userId int64) *errors_utils.RestError {
	user := &users.User{Id: userId}
	return user.Delete()
}
