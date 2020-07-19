package payment_dto

import (
	"fmt"
	"github.com/golanshy/plime_core-go/data_models/transaction_dto"
	"github.com/golanshy/plime_core-go/data_models/user_dto"
	"github.com/golanshy/plime_core-go/data_models/wallet_dao"
	"github.com/golanshy/plime_core-go/utils/date_utils"
	"github.com/golanshy/plime_core-go/utils/rest_errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"strings"
	"time"
)

type PaymentsRequest struct {
	Id          string    `json:"id"`
	Reference   string    `json:"reference,omitempty"`
	Details     string    `json:"details,omitempty"`
	Payments    []Payment `json:"payments,omitempty"`
	WebHook     WebHook   `json:"web_hook,omitempty"`
	DateCreated time.Time `json:"date_created,omitempty"`
}

type Payment struct {
	Id           string        `json:"id,omitempty"`
	Payee        user_dto.User `json:"payee,omitempty"`
	UserSecrets  []UserSecret  `json:"user_secrets,omitempty"`
	Reference    string        `json:"reference,omitempty"`
	Details      string        `json:"details,omitempty"`
	Amount       float64       `json:"amount,omitempty"`
	CurrencyCode string        `json:"currency_code,omitempty"`
	SendOn       string        `json:"send_on,omitempty"`
	ArriveBy     string        `json:"arrive_by,omitempty"`
	DateCreated  time.Time     `json:"date_created,omitempty"`
}

func (payment *Payment) Validate() *rest_errors.RestErr {
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
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}

type UserSecret struct {
	Key   string `json:"key"`
	Human string `json:"human,omitempty"`
	Value string `json:"value,omitempty"`
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
	ReferenceId interface{} `json:"reference_id,omitempty"`
	Reference   string      `json:"reference,omitempty"`
	Details     string      `json:"details,omitempty"`
	DateCreated *time.Time  `json:"date_created,omitempty"`
}

type PaymentResult struct {
	Id                 *primitive.ObjectID                  `json:"id,omitempty" bson:"_id, omitempty"`
	Payer              user_dto.User                        `json:"payer"`
	Payee              user_dto.User                        `json:"payee"`
	UserSecrets        *[]UserSecret                        `json:"user_secrets,omitempty"`
	WebHook            *WebHook                             `json:"web_hook,omitempty"`
	Reference          string                               `json:"reference,omitempty"`
	Details            string                               `json:"details,omitempty"`
	Amount             float64                              `json:"amount"`
	CurrencyCode       string                               `json:"currency_code"` // Iso 4217 https://en.wikipedia.org/wiki/ISO_4217
	SendOn             string                               `json:"send_on,omitempty"`
	ArriveBy           string                               `json:"arrive_by,omitempty"`
	Status             string                               `json:"status"`
	TransactionResults *[]transaction_dto.TransactionResult `json:"transaction_results,omitempty"`
	FailureDetails     string                               `json:"failure_details,omitempty"`
	Error              *rest_errors.RestErr                 `json:"error,omitempty"`
	DateCreated        string                               `json:"date_created,omitempty"`
	LastUpdated        string                               `json:"last_updated,omitempty"`
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
	if payeeErr := request.PayeeWallet.Validate(); payeeErr != nil {
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

type PaymentResultsResponse struct {
	Start   int64            `json:"start"`
	Limit   int64            `json:"limit"`
	Hits    int64            `json:"hits"`
	Total   int64            `json:"total"`
	Results *[]PaymentResult `json:"results,omitempty"`
}

type PaymentProcessRequest struct {
	Id          primitive.ObjectID `json:"id,omitempty" bson:"_id, omitempty"`
	UserSecrets *[]UserSecret      `json:"user_secrets,omitempty"`
	DateCreated time.Time          `json:"date_created,omitempty"`
}

func (request *PaymentProcessRequest) Validate() *rest_errors.RestErr {
	if request.Id.IsZero() {
		return rest_errors.NewBadRequestError("invalid id")
	}
	if len(*request.UserSecrets) > 0 {
		for _, requestSecret := range *request.UserSecrets {
			if requestSecret.Key != "" && strings.TrimSpace(requestSecret.Value) == "" {
				return rest_errors.NewBadRequestError(fmt.Sprintf("invalid %s", requestSecret.Human))
			}
		}
	}
	return nil
}
