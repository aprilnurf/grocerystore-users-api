package users

import (
	"github.com/aprilnurf/grocerystore_users-api/datasources/mysql"
	"github.com/aprilnurf/grocerystore_users-api/logger"
	"github.com/aprilnurf/grocerystore_users-api/utils/date_utils"
	"github.com/aprilnurf/grocerystore_users-api/utils/errors_utils"
)

const (
	queryInsert           = "INSERT INTO users(first_name, last_name, email, created_date, password, status) VALUES (?,?,?,?,?,?);"
	queryUpdate           = "UPDATE users SET first_name=?, last_name=?, email=? WHERE id=?;"
	queryGetUser          = "SELECT id, first_name, last_name, email, created_date, status FROM users WHERE id=?;"
	queryDeleteUser       = "DELETE FROM users WHERE id=?;"
	queryFindUserByStatus = "SELECT id, first_name, last_name, email, created_date, status FROM users WHERE status=?;"
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
		logger.Error("Error when trying to prepare user statement", err)
		return errors_utils.NewInternalServerError("database error")
	}
	defer stmt.Close()

	result := stmt.QueryRow(&user.Id)
	if err := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.CreatedDate, &user.Status); err != nil {
		logger.Error("Error when trying to query get user", err)
		return errors_utils.NewInternalServerError("database error")


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
		logger.Error("Error when trying to prepare insert user statement", err)
		return errors_utils.NewInternalServerError("database error")
	}
	defer rows.Close()

	result, saveErr := rows.Exec(user.FirstName, user.LastName, user.Email, user.CreatedDate, user.Password, user.Status)

	if saveErr != nil {
		logger.Error("Error when trying to save user", err)
		return errors_utils.NewInternalServerError("database error")
		//if strings.Contains(saveErr.Error(), indexUniqueEmail) {
		//	return errors.NewBadRequestError(fmt.Sprintf("Email %s already exist", user.Email))
		//}
		//return errors.NewInternalServerError(sqlErr.Error())
	}
	userId, err := result.LastInsertId()

	if err != nil {
		logger.Error("Error when trying to get last insert id after creating a new user", err)
		return errors_utils.NewInternalServerError("database error")
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
		logger.Error("Error when trying prepare update user statement", err)
		return errors_utils.NewInternalServerError("database error")
	}
	defer rows.Close()

	user.CreatedDate = date_utils.GetNowString()
	_, saveErr := rows.Exec(user.FirstName, user.LastName, user.Email, user.Id)

	if saveErr != nil {
		logger.Error("Error when trying to update user ", err)
		return errors_utils.NewInternalServerError("database error")
		//if strings.Contains(saveErr.Error(), indexUniqueEmail) {
		//	return errors.NewBadRequestError(fmt.Sprintf("Email %s already exist", user.Email))
		//}
		//return errors.NewInternalServerError(sqlErr.Error())
	}
	return nil
}

func (user *User) Delete() *errors_utils.RestError {
	stmt, err := mysql.Client.Prepare(queryDeleteUser)
	if err != nil {
		logger.Error("Error when trying to prepare delete user statement", err)
		return errors_utils.NewInternalServerError("database error")
	}
	defer stmt.Close()

	if _, err = stmt.Exec(user.Id); err != nil {
		logger.Error("Error when trying to delete user", err)
		return errors_utils.NewInternalServerError("database error")
	}
	return nil
}

func (user *User) FindByStatus(status bool) ([]User, *errors_utils.RestError) {
	stmt, err := mysql.Client.Prepare(queryFindUserByStatus)
	if err != nil {
		logger.Error("Error when trying to prepare find user statement", err)
		return nil, errors_utils.NewInternalServerError("database error")
	}
	defer stmt.Close()

	rows, err := stmt.Query(status)
	if err != nil {
		logger.Error("Error when trying to find user", err)
		return nil, errors_utils.NewInternalServerError("database error")
	}
	defer rows.Close()

	results := make([]User, 0)
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.CreatedDate, &user.Status); err != nil {
			logger.Error("Error when trying to scan find user", err)
			return nil, errors_utils.NewInternalServerError("database error")
		}
		results = append(results, user)
	}

	if len(results) == 0 {
		return nil, errors_utils.NewNotExistError("no users matching status")
	}
	return results, nil
}
