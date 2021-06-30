package mysql_utils

import (
	"strings"

	"github.com/go-sql-driver/mysql"
	"github.com/jgmc3012/bookstore_users-api/utils/errors"
)

const (
	errorNoRows = "no rows in result set"
)

func ParseError(err error) *errors.RestErr {
	slqErr, ok := err.(*mysql.MySQLError)
	if !ok {
		if strings.Contains(err.Error(), errorNoRows) {
			return errors.NewNotFoundError("no record matching given id")
		}
		return errors.NewInternalServerError("error parsing database response")
	}

	switch slqErr.Number {
	case 1062:
		return errors.NewBadRequestError("invalid data, duplicate field")
	}

	return errors.NewInternalServerError("error processing request")
}
