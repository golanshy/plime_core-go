package account_dto

import (
	"github.com/golanshy/plime_core-go/data_models/address_dto"
	"github.com/golanshy/plime_core-go/data_models/kyc_dto"
	"github.com/golanshy/plime_core-go/data_models/payment_dto"
	"github.com/golanshy/plime_core-go/data_models/user_dto"
	"github.com/golanshy/plime_core-go/data_models/wallet_dao"
	"github.com/golanshy/plime_core-go/utils/rest_errors"
	"strings"
)

type AccountRequest struct {
	Email        string              `json:"email"`
	AccountName  string              `json:"account_name"`
	AccountType  int                 `json:"account_type"`
	CurrencyCode string              `json:"currency_code"` // Iso 4217 https://en.wikipedia.org/wiki/ISO_4217
	Owner        user_dto.User       `json:"owner"`
	Address      address_dto.Address `json:"address"`
}

func (accountRequest *AccountRequest) Trim() {
	accountRequest.Email = strings.TrimSpace(accountRequest.Email)
	accountRequest.AccountName = strings.TrimSpace(accountRequest.AccountName)
	accountRequest.Owner.Trim()
	accountRequest.Address.Trim()
}

func (accountRequest *AccountRequest) Validate() *rest_errors.RestErr {
	accountRequest.Trim()
	if accountRequest.Email == "" {
		return rest_errors.NewBadRequestError("invalid email address")
	}
	return nil
}

type Account struct {
	Id            string               `json:"id"`
	Email         string               `json:"email,omitempty"`
	AccountName   string               `json:"account_name,omitempty"`
	AccountType   int                  `json:"account_type,omitempty"`
	Active        bool                 `json:"active,omitempty"`
	KycStatus     *kyc_dto.KycStatus   `json:"kyc_status,omitempty"`
	Details       string               `json:"details,omitempty"`
	Notes         string               `json:"notes,omitempty"`
	Owner         user_dto.User        `json:"owner,omitempty"`
	Users         []AccountUser        `json:"users,omitempty"`
	Address       address_dto.Address  `json:"address,omitempty"`
	Wallets       []wallet_dao.Wallet  `json:"wallets,omitempty"`
	Beneficiaries []AccountBeneficiary `json:"beneficiaries,omitempty"`
}

func (account *Account) Validate() *rest_errors.RestErr {
	account.Email = strings.TrimSpace(account.Email)
	account.AccountName = strings.TrimSpace(account.AccountName)
	if account.Email == "" {
		return rest_errors.NewBadRequestError("invalid email address")
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
