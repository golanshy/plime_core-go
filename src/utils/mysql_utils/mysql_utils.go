package mysql_utils

import (
	"github.com/go-sql-driver/mysql"
	"github.com/golanshy/plime_core-go/src/utils/rest_errors"
	"strings"
)

const (
	ErrorNoRows = "no rows in result set"
)

func ParseError(err error) rest_errors.RestErr {
	sqErr, ok := err.(*mysql.MySQLError)
	if !ok {
		if strings.Contains(err.Error(), ErrorNoRows) {
			return rest_errors.NewNotFoundError("No record matching given id")
		}
		return rest_errors.NewInternalServerError("error parsing database response", err)
	}
	switch sqErr.Number {
	case 1062:
		return rest_errors.NewBadRequestError("duplicate field")
	}
	return rest_errors.NewInternalServerError("error processing request", err)
}
