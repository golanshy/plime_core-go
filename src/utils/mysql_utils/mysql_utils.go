package mysql_utils

import (
	"github.com/go-sql-driver/mysql"
	"github.com/golanshy/plime_core-go/src/utils/errors"
	"strings"
)

const (
	ErrorNoRows = "no rows in result set"
)

func ParseError(err error) *errors.RestErr {
	sqErr, ok := err.(*mysql.MySQLError)
	if !ok {
		if strings.Contains(err.Error(), ErrorNoRows) {
			return errors.NewNotFoundError("No record matching given id")
		}
		return errors.NewInternalServerError("error parsing database response")
	}
	switch sqErr.Number {
	case 1062:
		return errors.NewBadRequestError("duplicate field")
	}
	return errors.NewInternalServerError("error processing request")
}
