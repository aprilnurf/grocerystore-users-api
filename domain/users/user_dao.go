package users

import (
	"github.com/aprilnurf/grocerystore_users-api/datasources/mysql"
	"github.com/aprilnurf/grocerystore_users-api/utils/date_utils"
	"github.com/aprilnurf/grocerystore_users-api/utils/errors_utils"
	"github.com/aprilnurf/grocerystore_users-api/utils/mysql_utils"
)

const (
	indexUniqueEmail = "email_UNIQUE"
	errorNoRows      = "no rows in result set"
	queryInsert      = "INSERT INTO users(first_name, last_name, email, created_date) VALUES (?,?,?,?);"
	queryUpdate      = "UPDATE users SET first_name=?, last_name=?, email=? WHERE id=?;"
	queryGetUser     = "SELECT id, first_name, last_name, email, created_date FROM users WHERE id=?"
)

var (
	usersDB = make(map[int64]*User)
)

func (user *User) Get() *errors_utils.RestError {
	//if err := mysql.Client.Ping(); err != nil {
	//	panic(err)
	//}
	//result := usersDB[user.Id]
	//if result == nil {
	//	return errors.NewNotExistError(fmt.Sprintf("User %d Not Exist", user.Id))
	//}
	//user.Id = result.Id
	//user.FirstName = result.FirstName
	//user.LastName = result.LastName
	//user.Email = result.Email
	//user.CreatedDate = result.CreatedDate
	//return nil
	stmt, err := mysql.Client.Prepare(queryGetUser)
	if err != nil {
		return errors_utils.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	result := stmt.QueryRow(&user.Id)
	if err := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.CreatedDate); err != nil {
		return mysql_utils.ParseError(err)
		//if strings.Contains(err.Error(), errorNoRows) {
		//	return errors_utils.NewNotExistError(fmt.Sprintf("user %d not found", user.Id))
		//
		//}
		//return errors_utils.NewInternalServerError(fmt.Sprintf("Error trying to get user %d", user.Id, err.Error()))
	}

	return nil
}

func (user *User) Save() *errors_utils.RestError {
	rows, err := mysql.Client.Prepare(queryInsert)
	if err != nil {
		return errors_utils.NewInternalServerError(err.Error())
	}
	defer rows.Close()

	user.CreatedDate = date_utils.GetNowString()
	result, saveErr := rows.Exec(user.FirstName, user.LastName, user.Email, user.CreatedDate)

	if saveErr != nil {
		return mysql_utils.ParseError(saveErr);
		//if strings.Contains(saveErr.Error(), indexUniqueEmail) {
		//	return errors.NewBadRequestError(fmt.Sprintf("Email %s already exist", user.Email))
		//}
		//return errors.NewInternalServerError(sqlErr.Error())
	}
	userId, err := result.LastInsertId()

	if err != nil {
		return errors_utils.NewInternalServerError(err.Error())
	}
	user.Id = userId
	//current := usersDB[user.Id]
	//if current != nil {
	//	if current.Email == user.Email {
	//		return errors.NewBadRequestError(fmt.Sprintf("email %s already register", user.Email))
	//	}
	//	return errors.NewBadRequestError(fmt.Sprintf("User %d already exist", user.Id))
	//}
	//usersDB[user.Id] = user
	return nil
}

func (user *User) Update() *errors_utils.RestError {
	rows, err := mysql.Client.Prepare(queryUpdate)
	if err != nil {
		return errors_utils.NewInternalServerError(err.Error())
	}
	defer rows.Close()

	user.CreatedDate = date_utils.GetNowString()
	_, saveErr := rows.Exec(user.FirstName, user.LastName, user.Email, user.Id)

	if saveErr != nil {
		return mysql_utils.ParseError(saveErr)
		//if strings.Contains(saveErr.Error(), indexUniqueEmail) {
		//	return errors.NewBadRequestError(fmt.Sprintf("Email %s already exist", user.Email))
		//}
		//return errors.NewInternalServerError(sqlErr.Error())
	}
	return nil
}
