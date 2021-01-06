package mysql_utils

import (
	"strings"

	"github.com/go-sql-driver/mysql"

	"smartshader/go-bookstore/users/utils/errors"
)

const (
	ErrMysqlDuplicateData = 1062
	errorNoRows           = "no rows in result set"
)

func ParseError(err error) *errors.RestErr {
	sqlErr, ok := err.(*mysql.MySQLError)
	if !ok {
		if strings.Contains(err.Error(), errorNoRows) {
			return errors.NewNotFoundError("no record matching given id")
		}

		return errors.NewInternalServerError("error parsing database response")
	}

	switch sqlErr.Number {
	case ErrMysqlDuplicateData:
		return errors.NewBadRequestError("invalid data")
	}

	return errors.NewInternalServerError("error processing request")
}
