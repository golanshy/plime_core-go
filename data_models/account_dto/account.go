package account_dto

import (
	"fmt"
	"github.com/golanshy/plime_core-go/data_models/address_dto"
	"github.com/golanshy/plime_core-go/data_models/customer_dto"
	"github.com/golanshy/plime_core-go/data_models/kyb_dto"
	"github.com/golanshy/plime_core-go/data_models/kyc_dto"
	"github.com/golanshy/plime_core-go/data_models/payment_dto"
	"github.com/golanshy/plime_core-go/data_models/user_dto"
	"github.com/golanshy/plime_core-go/data_models/wallet_dao"
	"github.com/golanshy/plime_core-go/utils/rest_errors"
	"strings"
)

//type PersonalAccount = 0
//type SoleTraderAccount = 1
//type BusinessAccount = 2

type AccountRequest struct {
	Type         int64                 `json:"type"`
	Name         string                `json:"name"`
	Email        string                `json:"email"`
	CurrencyCode string                `json:"currency_code"` // Iso 4217 https://en.wikipedia.org/wiki/ISO_4217
	Customer     customer_dto.Customer `json:"customer,omitempty"`
	Address      address_dto.Address   `json:"address"`
}

func (accountRequest *AccountRequest) Trim() {
	accountRequest.Email = strings.TrimSpace(accountRequest.Email)
	accountRequest.Name = strings.TrimSpace(accountRequest.Name)
	accountRequest.CurrencyCode = strings.TrimSpace(accountRequest.CurrencyCode)
	accountRequest.Customer.Trim()
	accountRequest.Address.Trim()
}

func (accountRequest *AccountRequest) Validate() *rest_errors.RestErr {
	accountRequest.Trim()
	if accountRequest.Email == "" {
		return rest_errors.NewBadRequestError("invalid email address field")
	}
	if accountRequest.Type < 0 || accountRequest.Type > 2 {
		return rest_errors.NewBadRequestError("invalid type field")
	}
	if accountRequest.Customer.Id == "" {
		return rest_errors.NewBadRequestError("invalid customer id field")
	}
	if accountRequest.Customer.Type != accountRequest.Type {
		return rest_errors.NewBadRequestError("invalid customer type field")
	}
	if accountRequest.Customer.Type < 0 || accountRequest.Customer.Type > 2 {
		return rest_errors.NewBadRequestError("invalid customer type field")
	}
	if accountRequest.Customer.Name == "" {
		return rest_errors.NewBadRequestError("invalid customer name field")
	}
	// For business accounts only
	if accountRequest.Type == 2 {
		if accountRequest.Customer.CompanyRegisteredName == "" {
			return rest_errors.NewBadRequestError("invalid customer company name field")
		}
		if accountRequest.Customer.CompanyRegisteredId == "" {
			return rest_errors.NewBadRequestError("invalid customer company id field")
		}
	}
	if err := accountRequest.Customer.Validate(); err != nil {
		return rest_errors.NewBadRequestError(fmt.Sprintf("invalid customer address details - %s", err.Message))
	}
	return nil
}

type AccountsResult struct {
	Start   int64     `json:"start"`
	Limit   int64     `json:"limit"`
	Hits    int64     `json:"hits"`
	Total   int64     `json:"total"`
	Results []Account `json:"results,omitempty"`
}

type Account struct {
	Id            string                 `json:"id"`
	Email         string                 `json:"email,omitempty"`
	Name          string                 `json:"name,omitempty"`
	Type          int64                  `json:"type,omitempty"`
	Suspended     bool                   `json:"suspended,omitempty"`
	KycStatus     *kyc_dto.KycStatus     `json:"kyc_status,omitempty"`
	KybStatus     *kyb_dto.KybStatus     `json:"kyb_status,omitempty"`
	Details       string                 `json:"details,omitempty"`
	Notes         string                 `json:"notes,omitempty"`
	Owner         *user_dto.User         `json:"owner,omitempty"`
	Customer      *customer_dto.Customer `json:"customer,omitempty"`
	Users         []AccountUser          `json:"users,omitempty"`
	Address       address_dto.Address    `json:"address,omitempty"`
	Wallets       []wallet_dao.Wallet    `json:"wallets,omitempty"`
	Beneficiaries []AccountBeneficiary   `json:"beneficiaries,omitempty"`
}

func (account *Account) Validate() *rest_errors.RestErr {
	account.Email = strings.TrimSpace(account.Email)
	account.Name = strings.TrimSpace(account.Name)
	if account.Email == "" {
		return rest_errors.NewBadRequestError("invalid email address field")
	}
	return nil
}

// AccountType:
// 0 = Personal
// 1 = Self Employed
// 2 = Company

type AccountUser struct {
	User user_dto.User `json:"user"`
	Role int           `json:"role"`
}

// Role:
// 0 = Viewer
// 1 = Owner
// 2 = Editor

type AccountExchangeRequest struct {
	Account          Account `json:"payer_account,omitempty"`
	Reference        string  `json:"reference,omitempty"`
	Details          string  `json:"details,omitempty"`
	Amount           float64 `json:"amount,omitempty"`
	FromCurrencyCode string  `json:"from_currency_code,omitempty"`
	ToCurrencyCode   string  `json:"to_currency_code,omitempty"`
	ExchangeOn       string  `json:"exchange_on,omitempty"`
}

func (request *AccountExchangeRequest) Validate() *rest_errors.RestErr {
	return nil
}

type AccountBeneficiary struct {
	User    user_dto.User `json:"user"`
	Account Account       `json:"payer_account"`
}

type AccountBeneficiaryPaymentRequest struct {
	PayerAccount       Account                          `json:"payer_account,omitempty"`
	AccountBeneficiary AccountBeneficiary               `json:"account_beneficiary,omitempty"`
	Payment            payment_dto.PublicPaymentRequest `json:"payment,omitempty"`
}

func (request *AccountBeneficiaryPaymentRequest) Validate() *rest_errors.RestErr {
	return request.Payment.Validate()
}
