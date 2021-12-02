package mysql_utils

import (
	"strings"

	"github.com/go-sql-driver/mysql"
	"github.com/jmillandev/bookstore_utils-go/rest_errors"
)

const (
	ErrorNoRows = "no rows in result set"
)

func ParseError(err error) *rest_errors.RestErr {
	slqErr, ok := err.(*mysql.MySQLError)
	if !ok {
		if strings.Contains(err.Error(), ErrorNoRows) {
			return rest_errors.NewNotFoundError("no record matching given id")
		}
		return rest_errors.NewInternalServerError("error parsing database response", err)
	}

	switch slqErr.Number {
	case 1062:
		return rest_errors.NewBadRequestError("invalid data, duplicate field")
	}

	return rest_errors.NewInternalServerError("error processing request", err)
}
