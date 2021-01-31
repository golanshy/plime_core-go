package customer_dto

import (
	"fmt"
	"github.com/golanshy/plime_core-go/data_models/address_dto"
	"github.com/golanshy/plime_core-go/data_models/user_dto"
	"github.com/golanshy/plime_core-go/utils/rest_errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"strings"
)

//type PersonalAccount = 0
//type SoleTraderAccount = 1
//type BusinessAccount = 2

type CustomersResult struct {
	Start   int64      `json:"start"`
	Limit   int64      `json:"limit"`
	Hits    int64      `json:"hits"`
	Total   int64      `json:"total"`
	Results []Customer `json:"results,omitempty"`
}

type Customer struct {
	Id                    primitive.ObjectID  `json:"id,omitempty" bson:"_id, omitempty"`
	BankingId             string              `json:"banking_id,omitempty"`
	Type                  int64               `json:"type"`
	Name                  string              `json:"name,omitempty"`
	Details               string              `json:"details,omitempty"`
	CompanyRegisteredName string              `json:"company_registered_name,omitempty"`
	CompanyRegisteredId   string              `json:"company_registered_id,omitempty"`
	ContactPerson         user_dto.User       `json:"contact_person,omitempty"`
	Status                string              `json:"status,omitempty"`
	Address               address_dto.Address `json:"address,omitempty"`
	CustomerUsers         []CustomerUser      `json:"customer_users,omitempty"`
	FeatureFlags          []string            `json:"feature_flags,omitempty"`
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

func (customer *Customer) Validate() *rest_errors.RestErr {
	customer.Trim()
	if customer.Type < 0 || customer.Type > 2 {
		return rest_errors.NewBadRequestError("invalid type field")
	}
	if customer.Name == "" {
		return rest_errors.NewBadRequestError("invalid name field")
	}
	if customer.Type == 2 {
		// Mandatory for Business Customers
		if customer.CompanyRegisteredName == "" {
			return rest_errors.NewBadRequestError("invalid company name field")
		}
		if customer.CompanyRegisteredId == "" {
			return rest_errors.NewBadRequestError("invalid company id field")
		}
	}

	if err := customer.Address.Validate(); err != nil {
		return rest_errors.NewBadRequestError(fmt.Sprintf("invalid customer address details - %s", err.Message))
	}
	return nil
}

type CustomerUser struct {
	User user_dto.User `json:"user,omitempty"`
	Role string        `json:"role,omitempty"`
}

func (user *CustomerUser) Trim() {
	user.User.Trim()
	user.Role = strings.TrimSpace(user.Role)
}
