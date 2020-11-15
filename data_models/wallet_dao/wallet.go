package wallet_dao

import (
	"github.com/golanshy/plime_core-go/logger"
	"github.com/golanshy/plime_core-go/utils/rest_errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"strings"
	"time"
)

type Wallet struct {
	Id               string            `json:"id"`
	HolderId         string            `json:"holder_id"`
	PartnerRef       string            `json:"partner_ref,omitempty"`
	PartnerProduct   string            `json:"partner_product,omitempty"`
	CurrencyCode     string            `json:"currency_code"` // Iso 4217 https://en.wikipedia.org/wiki/ISO_4217
	Iban             string            `json:"iban,omitempty"`
	BicSwift         string            `json:"bic_swift,omitempty"`
	WalletStatus     string            `json:"status"`
	IbanStatus       string            `json:"iban_status,omitempty"`
	Amount           float64           `json:"amount"`
	CreatedAt        time.Time         `json:"created_at,omitempty"`
	CountryCode      string            `json:"country_code,omitempty"`
	UkSortCode       string            `json:"uk_sort_code,omitempty"`
	UkAccountNumber  string            `json:"uk_account_number,omitempty"`
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

type RestrictedWallet struct {
	Id               primitive.ObjectID `json:"id,omitempty" bson:"_id, omitempty"`
	HolderId         string            `json:"holder_id"`
	WalletStatus     string            `json:"status"`
	CurrencyCode     string            `json:"currency_code"` // Iso 4217 https://en.wikipedia.org/wiki/ISO_4217
	Amount           float64           `json:"amount"`
	RestrictedTo     string             `json:"restricted_to"`
	PaymentReference string             `json:"payment_reference"`
	PayerId          string             `json:"payer_id"`
	PayerName        string             `json:"payer_name"`
	CreatedAt        time.Time         `json:"created_at,omitempty"`
}

type RestrictedWalletRequest struct {
	HolderId         string  `json:"holder_id"`
	RestrictedTo     string  `json:"restricted_to"`
	PaymentReference string  `json:"payment_reference"`
	PayerId          string  `json:"payer_id"`
	PayerName        string  `json:"payer_name"`
	CurrencyCode     string  `json:"currency_code"`
	Amount           float64 `json:"amount"`
}

func (request *RestrictedWalletRequest) Validate() *rest_errors.RestErr {
	if strings.TrimSpace(request.HolderId) == "" {
		logger.Error("error creating restricted wallet missing account id", nil)
		return rest_errors.NewBadRequestError("error creating restricted wallet missing account id")
	}
	if strings.TrimSpace(request.RestrictedTo) == "" {
		logger.Error("error creating restricted wallet missing restrictedTo code", nil)
		return rest_errors.NewBadRequestError("error creating restricted wallet missing restrictedTo code")
	}
	if strings.TrimSpace(request.PaymentReference) == "" {
		logger.Error("error creating restricted wallet missing payment reference", nil)
		return rest_errors.NewBadRequestError("error creating restricted wallet missing payment reference")
	}
	if strings.TrimSpace(request.PayerId) == "" {
		logger.Error("error creating restricted wallet missing payerId", nil)
		return rest_errors.NewBadRequestError("error creating restricted wallet missing payerId")
	}
	if strings.TrimSpace(request.PayerName) == "" {
		logger.Error("error creating restricted wallet missing payerName", nil)
		return rest_errors.NewBadRequestError("error creating restricted wallet missing payerName")
	}
	if request.Amount <= 0 {
		logger.Error("error creating restricted wallet invalid amount", nil)
		return rest_errors.NewBadRequestError("error creating restricted wallet invalid amount")
	}
	if strings.TrimSpace(request.CurrencyCode) == "" {
		logger.Error("error creating restricted wallet missing currency code", nil)
		return rest_errors.NewBadRequestError("error creating restricted wallet missing currency code")
	}
	return nil
}
