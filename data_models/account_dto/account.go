package account_dto

import (
	"github.com/golanshy/plime_core-go/data_models/address_dto"
	"github.com/golanshy/plime_core-go/data_models/user_dto"
	"github.com/golanshy/plime_core-go/data_models/wallet_dao"
	"github.com/golanshy/plime_core-go/utils/rest_errors"
	"strings"
)

type Account struct {
	Id             string              `json:"id"`
	Email          string              `json:"email"`
	AccountName    string              `json:"account_name"`
	AccountType    int                 `json:"account_type"`
	OwnerFirstName string              `json:"owner_first_name"`
	OwnerLastName  string              `json:"owner_last_name"`
	Address        address_dto.Address `json:"address"`
	Wallets        []wallet_dao.Wallet `json:"wallets"`
}

func (account *Account) Validate() *rest_errors.RestErr {
	account.Email = strings.TrimSpace(account.Email)
	account.AccountName = strings.TrimSpace(account.AccountName)
	account.OwnerFirstName = strings.TrimSpace(account.OwnerFirstName)
	account.OwnerLastName = strings.TrimSpace(account.OwnerLastName)
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