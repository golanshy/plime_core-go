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

type AccountRequest struct {
	Email        string                        `json:"email"`
	AccountName  string                        `json:"account_name"`
	AccountType  int                           `json:"account_type"`
	CurrencyCode string                        `json:"currency_code"` // Iso 4217 https://en.wikipedia.org/wiki/ISO_4217
	Owner        *user_dto.User                `json:"owner,omitempty"`
	Customer     *customer_dto.Customer `json:"customer,omitempty"`
	Address      address_dto.Address           `json:"address"`
}

func (accountRequest *AccountRequest) Trim() {
	accountRequest.Email = strings.TrimSpace(accountRequest.Email)
	accountRequest.AccountName = strings.TrimSpace(accountRequest.AccountName)
	if accountRequest.Owner != nil {
		accountRequest.Owner.Trim()
	}
	if accountRequest.Customer != nil {
		accountRequest.Customer.Trim()
	}
	accountRequest.Address.Trim()
}

func (accountRequest *AccountRequest) Validate() *rest_errors.RestErr {
	accountRequest.Trim()
	if accountRequest.Email == "" {
		return rest_errors.NewBadRequestError("invalid email address field")
	}
	if accountRequest.Owner != nil && accountRequest.Customer != nil {
		return rest_errors.NewBadRequestError("invalid owner and customer fields, cannot process both")
	}
	if accountRequest.Customer.Name == "" {
		return rest_errors.NewBadRequestError("invalid name field")
	}
	if accountRequest.Customer.CompanyRegisteredName == "" {
		return rest_errors.NewBadRequestError("invalid company name field")
	}
	if accountRequest.Customer.CompanyRegisteredId == "" {
		return rest_errors.NewBadRequestError("invalid company id field")
	}
	if err := accountRequest.Customer.Address.Validate(); err != nil {
		return rest_errors.NewBadRequestError(fmt.Sprintf("invalid customer address details - %s", err.Message))
	}
	return nil
}

type Account struct {
	Id            string                 `json:"id"`
	Email         string                 `json:"email,omitempty"`
	AccountName   string                 `json:"account_name,omitempty"`
	AccountType   int                    `json:"account_type,omitempty"`
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
	account.AccountName = strings.TrimSpace(account.AccountName)
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
	PayerAccount       Account             `json:"payer_account,omitempty"`
	AccountBeneficiary AccountBeneficiary  `json:"account_beneficiary,omitempty"`
	Payment            payment_dto.Payment `json:"payment,omitempty"`
}

func (request *AccountBeneficiaryPaymentRequest) Validate() *rest_errors.RestErr {
	return nil
}
