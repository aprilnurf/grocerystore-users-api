package users

type User struct {
	Id          int64 `json:"id"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	Email       string `json:"email"`
	CreatedDate string `json:"createdDate"`
}
