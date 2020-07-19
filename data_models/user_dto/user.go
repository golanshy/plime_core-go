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
	Id               int64                `json:"id"`
	FirstName        string               `json:"first_name,omitempty"`
	LastName         string               `json:"last_name,omitempty"`
	Email            string               `json:"email,omitempty"`
	DateOfBirth      string               `json:"date_of_birth,omitempty"`
	Mobile           string               `json:"mobile,omitempty"`
	CountryCode      string               `json:"country_code,omitempty"`
	DateCreated      string               `json:"date_created,omitempty"`
	Status           string               `json:"status,omitempty"`
	Passcode         string               `json:"passcode,omitempty"`
	BiometricEnabled bool                 `json:"biometric_enabled,omitempty"`
	EmailValidated   bool                 `json:"email_validated,omitempty"`
	MobileValidated  bool                 `json:"mobile_validated,omitempty"`
	Type             int64                `json:"type,omitempty"`
	Address          *address_dto.Address `json:"address,omitempty"`
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
