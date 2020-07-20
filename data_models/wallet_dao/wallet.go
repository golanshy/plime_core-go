package wallet_dao

import (
	"github.com/golanshy/plime_core-go/utils/rest_errors"
	"time"
)

type Wallet struct {
	Id              string     `json:"id"`
	HolderId        string     `json:"holder_id"`
	PartnerRef      string     `json:"partner_ref,omitempty"`
	PartnerProduct  string     `json:"partner_product,omitempty"`
	CurrencyCode    string     `json:"currency_code"` // Iso 4217 https://en.wikipedia.org/wiki/ISO_4217
	Iban            string     `json:"iban,omitempty"`
	BicSwift        string     `json:"bic_swift,omitempty"`
	WalletStatus    string     `json:"status"`
	IbanStatus      string     `json:"iban_status,omitempty"`
	Amount          float64    `json:"amount"`
	CreatedAt       time.Time `json:"created_at,omitempty"`
	CountryCode     string     `json:"country_code,omitempty"`
	UkSortCode      string     `json:"uk_sort_code,omitempty"`
	UkAccountNumber string     `json:"uk_account_number,omitempty"`
}

func (wallet *Wallet) Validate() *rest_errors.RestErr {
	if wallet.Id == "" {
		return rest_errors.NewBadRequestError("invalid wallet id")
	}
	if wallet.HolderId == "" {
		return rest_errors.NewBadRequestError("invalid wallet holder id")
	}
	if wallet.CurrencyCode == "" {
		return rest_errors.NewBadRequestError("invalid wallet currency code")
	}
	return nil
}
