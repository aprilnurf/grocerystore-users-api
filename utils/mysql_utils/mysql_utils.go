package mysql_utils

import (
	"github.com/aprilnurf/grocerystore_users-api/utils/errors_utils"
	"github.com/go-sql-driver/mysql"
	"strings"
)

func ParseError(err error) *errors_utils.RestError {
	sqlErr, ok := err.(*mysql.MySQLError)
	if !ok {
		if strings.Contains(err.Error(), "") {
			return errors_utils.NewNotExistError("no record matching given id")
		}
		return errors_utils.NewInternalServerError("error parsing database response")
	}

	switch sqlErr.Number {
	case 1062:
		return errors_utils.NewBadRequestError("invalid data")
	}
	return errors_utils.NewInternalServerError("error when trying save request")
}
