package account_dto

import (
	"github.com/golanshy/plime_core-go/data_models/address_dto"
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
	Email         string               `json:"email"`
	AccountName   string               `json:"account_name"`
	AccountType   int                  `json:"account_type"`
	Active        bool                 `json:"active"`
	Owner         user_dto.User        `json:"owner"`
	Users         []AccountUser        `json:"users"`
	Address       address_dto.Address  `json:"address"`
	Wallets       []wallet_dao.Wallet  `json:"wallets"`
	Beneficiaries []AccountBeneficiary `json:"beneficiaries"`
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
	Account          Account `json:"payer_account"`
	Reference        string  `json:"reference"`
	Details          string  `json:"details"`
	Amount           float64 `json:"amount"`
	FromCurrencyCode string  `json:"from_currency_code"`
	ToCurrencyCode   string  `json:"to_currency_code"`
	ExchangeOn       string  `json:"exchange_on"`
}

func (request *AccountExchangeRequest) Validate() *rest_errors.RestErr {
	return nil
}

type AccountBeneficiary struct {
	User    user_dto.User `json:"user"`
	Account Account       `json:"payer_account"`
}

type AccountBeneficiaryPaymentRequest struct {
	PayerAccount       Account             `json:"payer_account"`
	AccountBeneficiary AccountBeneficiary  `json:"account_beneficiary"`
	Payment            payment_dto.Payment `json:"payment"`
}

func (request *AccountBeneficiaryPaymentRequest) Validate() *rest_errors.RestErr {
	return nil
}
