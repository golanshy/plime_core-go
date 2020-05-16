package account_dto

import (
	"github.com/golanshy/plime_core-go/data_models/address_dto"
	"github.com/golanshy/plime_core-go/data_models/user_dto"
	"github.com/golanshy/plime_core-go/data_models/wallet_dao"
	"github.com/golanshy/plime_core-go/utils/rest_errors"
	"strings"
)

type Account struct {
	Id            string               `json:"id"`
	Email         string               `json:"email"`
	AccountName   string               `json:"account_name"`
	AccountType   int                  `json:"account_type"`
	Active        bool                 `json:"active"`
	Owner         user_dto.User        `json:"owner"`
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

type AccountBeneficiary struct {
	User user_dto.User `json:"user"`
	Role int           `json:"role"`
}

// Role:
// 0 = Viewer
// 1 = Owner
// 2 = Editor
