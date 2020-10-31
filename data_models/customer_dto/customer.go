package customer_dto

import (
	"fmt"
	"github.com/golanshy/plime_core-go/data_models/address_dto"
	"github.com/golanshy/plime_core-go/data_models/user_dto"
	"github.com/golanshy/plime_core-go/utils/rest_errors"
	"strings"
)

type CustomerRequest struct {
	Name                  string              `json:"name,omitempty"`
	Details               string              `json:"details,omitempty"`
	CompanyRegisteredName string              `json:"company_registered_name,omitempty"`
	CompanyRegisteredId   string              `json:"company_registered_id,omitempty"`
	Address               address_dto.Address `json:"address,omitempty"`
}

func (customerRequest *CustomerRequest) Trim() {
	customerRequest.Name = strings.TrimSpace(customerRequest.Name)
	customerRequest.Details = strings.TrimSpace(customerRequest.Details)
}

func (customerRequest *CustomerRequest) Validate() *rest_errors.RestErr {
	customerRequest.Trim()
	if customerRequest.Name == "" {
		return rest_errors.NewBadRequestError("invalid name field")
	}
	if customerRequest.CompanyRegisteredId == "" {
		return rest_errors.NewBadRequestError("invalid company id field")
	}
	if customerRequest.Details == "" {
		return rest_errors.NewBadRequestError("invalid details field")
	}
	if err := customerRequest.Address.Validate(); err != nil {
		return rest_errors.NewBadRequestError(fmt.Sprintf("invalid customer address details - %s", err.Message))
	}
	return nil
}

type Customer struct {
	Id                    string              `json:"id"`
	Name                  string              `json:"name,omitempty"`
	Details               string              `json:"details,omitempty"`
	CompanyRegisteredName string              `json:"company_registered_name,omitempty"`
	CompanyRegisteredId   string              `json:"company_registered_id,omitempty"`
	ContactPerson         user_dto.User       `json:"contact_person,omitempty"`
	Status                string              `json:"status,omitempty"`
	Address               address_dto.Address `json:"address,omitempty"`
	CustomerUsers         []CustomerUser      `json:"customer_users,omitempty"`
	DateCreated           string              `json:"date_created,omitempty"`
}

func (customer *Customer) Trim() {
	customer.Name = strings.TrimSpace(customer.Name)
	customer.Details = strings.TrimSpace(customer.Details)
	customer.CompanyRegisteredName = strings.TrimSpace(customer.CompanyRegisteredName)
	customer.CompanyRegisteredId = strings.TrimSpace(customer.CompanyRegisteredId)
	customer.ContactPerson.Trim()
	customer.Address.Trim()
	for _, user := range customer.CustomerUsers {
		user.Trim()
	}
}

type CustomerUser struct {
	User user_dto.User `json:"user,omitempty"`
	Role string        `json:"role,omitempty"`
}

func (user *CustomerUser) Trim() {
	user.User.Trim()
	user.Role = strings.TrimSpace(user.Role)
}
