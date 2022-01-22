package users

import "encoding/json"

type PublicUser struct {
	Id          int64  `json:"user_id"`
	CreatedDate string `json:"createdDate"`
	Status      bool `json:"status"`
}

type PrivateUser struct {
	Id          int64  `json:"id"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	Email       string `json:"email"`
	CreatedDate string `json:"createdDate"`
	Status      bool `json:"status"`
}

type Users []User

func (user *User) Marshall(isPublic bool) interface{} {
	if isPublic {
		return PublicUser{
			Id:          user.Id,
			CreatedDate: user.CreatedDate,
			Status:      user.Status,
		}
	}
	userJson , _ := json.Marshal(user)
	var privateUser PrivateUser
	json.Unmarshal(userJson, &privateUser)
	return privateUser
}

func (users Users) Marshall(isPublic bool) []interface{} {
	result := make([]interface{}, len(users))
	for index, user := range users {
		result[index] = user.Marshall(isPublic)
	}
	return result
}


