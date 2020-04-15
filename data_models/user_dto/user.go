package user_dto

import (
	"github.com/golanshy/plime_core-go/data_models/address_dto"
	"github.com/golanshy/plime_core-go/utils/rest_errors"
	"strings"
)

const (
	StatusActive = "active"
)

type User struct {
	Id              int64               `json:"id"`
	FirstName       string              `json:"first_name"`
	LastName        string              `json:"last_name"`
	Email           string              `json:"email"`
	DateOfBirth     string              `json:"date_of_birth"`
	Mobile          string              `json:"mobile"`
	CountryCode     string              `json:"country_code"`
	DateCreated     string              `json:"date_created"`
	Status          string              `json:"status"`
	Passcode        string              `json:"passcode"`
	EmailValidated  bool                `json:"email_validated"`
	MobileValidated bool                `json:"mobile_validated"`
	Address         address_dto.Address `json:"address"`
}

type Users []User

func (user *User) Trim() {
	user.FirstName = strings.TrimSpace(user.FirstName)
	user.LastName = strings.TrimSpace(user.LastName)
	user.Email = strings.TrimSpace(user.Email)
	user.CountryCode = strings.TrimSpace(user.CountryCode)
	user.Mobile = strings.TrimSpace(user.Mobile)
}

func (user *User) Validate() *rest_errors.RestErr {
	user.Trim()
	if user.Email == "" {
		return rest_errors.NewBadRequestError("invalid email address")
	}
	return nil
}