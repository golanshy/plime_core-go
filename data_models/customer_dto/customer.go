package customer_dto

import (
	"fmt"
	"github.com/golanshy/plime_core-go/data_models/address_dto"
	"github.com/golanshy/plime_core-go/data_models/user_dto"
	"github.com/golanshy/plime_core-go/utils/rest_errors"
	"strings"
)

type Customer struct {
	Id                    string              `json:"id"`
	Name                  string              `json:"name,omitempty"`
	Details               string              `json:"details,omitempty"`
	CompanyRegisteredName string              `json:"company_registered_name,omitempty"`
	CompanyRegisteredId   string              `json:"company_registered_id,omitempty"`
	ContactPerson         user_dto.User       `json:"contact_person,omitempty"`
	Status                string              `json:"status,omitempty"`
	Address               address_dto.Address `json:"address,omitempty"`
	DateCreated           string              `json:"date_created,omitempty"`
}

func (customer *Customer) Trim() {
	customer.Name = strings.TrimSpace(customer.Name)
	customer.Details = strings.TrimSpace(customer.Details)
}

func (customer *Customer) Validate() *rest_errors.RestErr {
	customer.Trim()
	if customer.Name == "" {
		return rest_errors.NewBadRequestError("invalid name field")
	}
	if customer.CompanyRegisteredId == "" {
		return rest_errors.NewBadRequestError("invalid company id field")
	}
	if customer.Details == "" {
		return rest_errors.NewBadRequestError("invalid details field")
	}
	if customer.ContactPerson.FirstName == "" {
		return rest_errors.NewBadRequestError("invalid contact person first name field")
	}
	if customer.ContactPerson.LastName == "" {
		return rest_errors.NewBadRequestError("invalid contact person last name field")
	}
	if customer.ContactPerson.Email == "" {
		return rest_errors.NewBadRequestError("invalid contact person field")
	}
	if err := customer.Address.Validate(); err != nil {
		return rest_errors.NewBadRequestError(fmt.Sprintf("invalid customer address details - %s", err.Message))
	}
	return nil
}
