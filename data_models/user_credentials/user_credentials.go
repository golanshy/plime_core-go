package user_credentials

import (
	"github.com/golanshy/plime_core-go/logger"
	"github.com/golanshy/plime_core-go/utils/rest_errors"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (request LoginRequest) Validate() *rest_errors.RestErr {
	if request.Email == "" {
		logger.Error("error when trying to validate join request invalid email", nil)
		return rest_errors.NewBadRequestError("invalid email")
	}
	if request.Password == "" {
		logger.Error("error when trying to validate join request invalid password", nil)
		return rest_errors.NewBadRequestError("invalid password")
	}
	return nil
}

type UserCredentials struct {
	Email       string `json:"email"`
	Password    string `json:"password"`
	DateCreated string `json:"date_created"`
}