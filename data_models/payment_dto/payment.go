package payment_dto

import (
	"fmt"
	"github.com/golanshy/plime_core-go/data_models/transaction_dto"
	"github.com/golanshy/plime_core-go/data_models/user_dto"
	"github.com/golanshy/plime_core-go/data_models/wallet_dao"
	"github.com/golanshy/plime_core-go/utils/date_utils"
	"github.com/golanshy/plime_core-go/utils/rest_errors"
	"strings"
	"time"
)

type PaymentsRequest struct {
	Id          string    `json:"id"`
	Reference   string    `json:"reference"`
	Details     string    `json:"details"`
	Payments    []Payment `json:"payments"`
	WebHook     WebHook   `json:"web_hook"`
	DateCreated string    `json:"date_created"`
}

type Payment struct {
	Id           string        `json:"id"`
	Payer        user_dto.User `json:"payer"`
	Payee        user_dto.User `json:"payee"`
	UserSecrets  []UserSecret  `json:"user_secrets"`
	Reference    string        `json:"reference"`
	Details      string        `json:"details"`
	Amount       float64       `json:"amount"`
	CurrencyCode string        `json:"currency_code"`
	SendOn       string        `json:"send_on"`
	ArriveBy     string        `json:"arrive_by"`
	DateCreated  string        `json:"date_created"`
}

func (payment *Payment) Validate() *rest_errors.RestErr {
	if userErr := payment.Payer.Validate(); userErr != nil {
		return rest_errors.NewBadRequestError("invalid payer data")
	}
	if userErr := payment.Payee.Validate(); userErr != nil {
		return rest_errors.NewBadRequestError("invalid payee data")
	}
	if payment.Amount <= 0 {
		return rest_errors.NewBadRequestError("invalid amount")
	}
	if strings.TrimSpace(payment.CurrencyCode) == "" {
		return rest_errors.NewBadRequestError("invalid currency code field")
	}
	if len(payment.UserSecrets) > 0 {
		for _, secret := range payment.UserSecrets {
			if strings.TrimSpace(secret.Key) == "" {
				return rest_errors.NewBadRequestError("invalid user secret key field")
			}
			if strings.TrimSpace(secret.Human) == "" {
				return rest_errors.NewBadRequestError("invalid user secret human field")
			}
			if strings.TrimSpace(secret.Value) == "" {
				return rest_errors.NewBadRequestError("invalid user secret value field")
			}
		}
	}
	return nil
}

type WebHook struct {
	URL      string `json:"url"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserSecret struct {
	Key   string `json:"key"`
	Human string `json:"human"`
	Value string `json:"value"`
}

func (request *PaymentsRequest) Validate() *rest_errors.RestErr {
	request.Reference = strings.TrimSpace(request.Reference)
	request.Details = strings.TrimSpace(request.Details)

	if request.Reference == "" {
		return rest_errors.NewBadRequestError("invalid reference data")
	}
	if request.Payments == nil || len(request.Payments) == 0 {
		return rest_errors.NewBadRequestError("invalid payments data")
	}

	for index, payment := range request.Payments {
		payment.Reference = strings.TrimSpace(payment.Reference)
		payment.Details = strings.TrimSpace(payment.Details)
		payment.ArriveBy = strings.TrimSpace(payment.ArriveBy)
		payment.SendOn = strings.TrimSpace(payment.SendOn)
		payment.CurrencyCode = strings.TrimSpace(payment.CurrencyCode)
		payment.Payee.Email = strings.TrimSpace(payment.Payee.Email)
		if payment.Payee.Email == "" {
			return rest_errors.NewBadRequestError(fmt.Sprintf("invalid payee email address for payment %d", index))
		}
		if payment.Amount <= 0 {
			return rest_errors.NewBadRequestError(fmt.Sprintf("invalid amount for payment %d", index))
		}
		if payment.CurrencyCode == "" {
			return rest_errors.NewBadRequestError(fmt.Sprintf("invalid currency code for payment %d", index))
		}
		if payment.ArriveBy != "" {
			arriveBy, timeFormatErr := date_utils.FormatAPITime(payment.ArriveBy)
			if timeFormatErr != nil {
				return rest_errors.NewBadRequestError(fmt.Sprintf("%s, invalid arrive by time for payment %d", timeFormatErr.Message, index))
			}
			if arriveBy.Before(time.Now().UTC()) {
				return rest_errors.NewBadRequestError(fmt.Sprintf("arrive by should be in the future for payment %d", index))
			}
		}
		if payment.SendOn != "" {
			sendOn, timeFormatErr := date_utils.FormatAPITime(payment.SendOn)
			if timeFormatErr != nil {
				return rest_errors.NewBadRequestError(fmt.Sprintf("%s, invalid send on for payment %d", timeFormatErr.Message, index))
			}
			if sendOn.Before(time.Now().UTC()) {
				return rest_errors.NewBadRequestError(fmt.Sprintf("send on should be in the future for payment %d", index))
			}
		}

		if payment.ArriveBy != "" && payment.SendOn != "" {
			arriveBy, _ := date_utils.FormatAPITime(payment.ArriveBy)
			sendOn, _ := date_utils.FormatAPITime(payment.SendOn)
			if arriveBy != nil && sendOn != nil {
				if !sendOn.UTC().Before(arriveBy.UTC()) {
					return rest_errors.NewBadRequestError(fmt.Sprintf("send on should be before arrive by for payment %d", index))
				}
			}
		}
	}

	return nil
}

type PaymentsResponse struct {
	Id          string      `json:"id"`
	ReferenceId interface{} `json:"reference_id"`
	Reference   string      `json:"reference"`
	Details     string      `json:"details"`
	DateCreated string      `json:"date_created"`
}

type PaymentResult struct {
	Id                string                            `json:"id"`
	User              user_dto.User                     `json:"user"`
	Reference         string                            `json:"reference"`
	Details           string                            `json:"details"`
	Amount            float64                           `json:"amount"`
	CurrencyCode      string                            `json:"currency_code"` // Iso 4217 https://en.wikipedia.org/wiki/ISO_4217
	SendOn            string                            `json:"send_on"`
	ArriveBy          string                            `json:"arrive_by"`
	Status            string                            `json:"status"`
	TransactionResult transaction_dto.TransactionResult `json:"transaction_result"`
	FailureDetails    string                            `json:"failure_details,omitempty"`
	Error             *rest_errors.RestErr              `json:"error,omitempty"`
	DateCreated       string                            `json:"date_created"`
}

func (request *PaymentsResponse) Validate() *rest_errors.RestErr {
	//if request.Email == "" {
	//	return rest_errors.NewBadRequestError("invalid email address")
	//}
	return nil
}

type WalletPaymentRequest struct {
	PayerWallet wallet_dao.Wallet `json:"payer_wallet"`
	PayeeWallet wallet_dao.Wallet `json:"payee_wallet"`
	Payment     Payment           `json:"payment"`
}

func (request *WalletPaymentRequest) Validate() *rest_errors.RestErr {
	if payerErr := request.PayerWallet.Validate(); payerErr != nil {
		return rest_errors.NewBadRequestError("invalid payer data")
	}
	if payeeErr := request.PayerWallet.Validate(); payeeErr != nil {
		return rest_errors.NewBadRequestError("invalid payee data")
	}
	if paymentErr := request.Payment.Validate(); paymentErr != nil {
		return rest_errors.NewBadRequestError("invalid payment data")
	}
	if request.Payment.CurrencyCode != request.PayerWallet.CurrencyCode {
		return rest_errors.NewBadRequestError("payer wallet currency code does not match request")
	}
	if request.Payment.CurrencyCode != request.PayeeWallet.CurrencyCode {
		return rest_errors.NewBadRequestError("payee wallet currency code does not match request")
	}
	return nil
}
