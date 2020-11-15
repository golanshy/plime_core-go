package payment_dto

import (
	"fmt"
	"github.com/golanshy/plime_core-go/data_models/customer_dto"
	"github.com/golanshy/plime_core-go/data_models/transaction_dto"
	"github.com/golanshy/plime_core-go/data_models/wallet_dao"
	"github.com/golanshy/plime_core-go/utils/date_utils"
	"github.com/golanshy/plime_core-go/utils/rest_errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"strings"
	"time"
)

type PublicPaymentsRequest struct {
	Reference string                 `json:"reference,omitempty"`
	Details   string                 `json:"details,omitempty"`
	Payments  []PublicPaymentRequest `json:"payments,omitempty"`
	WebHook   WebHook                `json:"web_hook,omitempty"`
}

type PaymentsRequest struct {
	Id          primitive.ObjectID     `json:"id,omitempty" bson:"_id, omitempty"`
	Reference   string                 `json:"reference,omitempty"`
	Details     string                 `json:"details,omitempty"`
	Payments    []PublicPaymentRequest `json:"payments,omitempty"`
	WebHook     WebHook                `json:"web_hook,omitempty"`
	DateCreated time.Time              `json:"date_created,omitempty"`
}

func (request *PublicPaymentsRequest) Validate() *rest_errors.RestErr {
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

type PublicPaymentRequest struct {
	Type         int64                  `json:"type"`
	Reference    string                 `json:"reference"`
	Details      string                 `json:"details"`
	Amount       float64                `json:"amount"`
	CurrencyCode string                 `json:"currency_code"`
	Payer        *customer_dto.Customer `json:"payer,omitempty"`
	Payee        Payee                  `json:"payee"`
	UserSecrets  []UserSecret           `json:"user_secrets,omitempty"`
	RestrictedTo string                 `json:"restricted_to,omitempty"`
	SendOn       string                 `json:"send_on,omitempty"`
	ArriveBy     string                 `json:"arrive_by,omitempty"`
	Status       string                 `json:"status,omitempty"`
	DateCreated  time.Time              `json:"date_created,omitempty"`
}

type PaymentRequest struct {
	Id                 primitive.ObjectID                   `json:"id,omitempty" bson:"_id"`
	Type               int64                                `json:"type"`
	Reference          string                               `json:"reference"`
	Details            string                               `json:"details"`
	Amount             float64                              `json:"amount"`
	CurrencyCode       string                               `json:"currency_code"`
	Payer              customer_dto.Customer                `json:"payer"`
	Payee              Payee                                `json:"payee"`
	UserSecrets        []UserSecret                         `json:"user_secrets,omitempty"`
	RestrictedTo       string                               `json:"restricted_to,omitempty"`
	SendOn             string                               `json:"send_on,omitempty"`
	ArriveBy           string                               `json:"arrive_by,omitempty"`
	WebHook            WebHook                              `json:"web_hook,omitempty"`
	Status             string                               `json:"status"`
	TransactionResults *[]transaction_dto.TransactionResult `json:"transaction_results,omitempty"`
	FailureDetails     string                               `json:"failure_details,omitempty"`
	Error              *rest_errors.RestErr                 `json:"error,omitempty"`
	DateCreated        time.Time                            `json:"date_created,omitempty"`
	LastUpdated        time.Time                            `json:"last_updated,omitempty"`
}

func (paymentRequest *PublicPaymentRequest) Validate() *rest_errors.RestErr {
	if userErr := paymentRequest.Payee.Validate(); userErr != nil {
		return rest_errors.NewBadRequestError("invalid payee data")
	}
	if paymentRequest.Amount <= 0 {
		return rest_errors.NewBadRequestError("invalid amount")
	}
	if strings.TrimSpace(paymentRequest.CurrencyCode) == "" {
		return rest_errors.NewBadRequestError("invalid currency code field")
	}
	if len(paymentRequest.UserSecrets) > 0 {
		for _, secret := range paymentRequest.UserSecrets {
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

type Payee struct {
	Type                int64  `json:"type"`
	Details             string `json:"details,omitempty"`
	CompanyName         string `json:"company_name,omitempty"`
	CompanyRegisteredId string `json:"company_registered_id,omitempty"`
	FirstName           string `json:"first_name,omitempty"`
	LastName            string `json:"last_name,omitempty"`
	Email               string `json:"email,omitempty"`
	Mobile              string `json:"mobile,omitempty"`
	CountryCode         string `json:"country_code,omitempty"`
}

func (payee *Payee) Trim() {
	payee.Details = strings.TrimSpace(payee.Details)
	payee.CompanyName = strings.TrimSpace(payee.CompanyName)
	payee.CompanyRegisteredId = strings.TrimSpace(payee.CompanyRegisteredId)
	payee.FirstName = strings.TrimSpace(payee.FirstName)
	payee.LastName = strings.TrimSpace(payee.LastName)
	payee.Email = strings.TrimSpace(payee.Email)
	payee.Mobile = strings.TrimSpace(payee.Mobile)
	payee.CountryCode = strings.TrimSpace(payee.CountryCode)
}

func (payee *Payee) Validate() *rest_errors.RestErr {
	payee.Trim()
	if payee.Email == "" {
		return rest_errors.NewBadRequestError("invalid payee email")
	}
	if payee.Type < 0 || payee.Type > 2 {
		return rest_errors.NewBadRequestError("invalid payee type")
	}
	// business type
	if payee.Type == 2 && (payee.CompanyName == "" || payee.CompanyRegisteredId == "") {
		return rest_errors.NewBadRequestError("invalid company details for business type")
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

type PaymentResultsResponse struct {
	Start   int64            `json:"start"`
	Limit   int64            `json:"limit"`
	Hits    int64            `json:"hits"`
	Total   int64            `json:"total"`
	Results []PaymentRequest `json:"results,omitempty"`
}

type PaymentProcessRequest struct {
	Id                string       `json:"id"`
	ThirdPartyPayeeId string       `json:"third_party_payee_id,omitempty"`
	UserSecrets       []UserSecret `json:"user_secrets,omitempty"`
	DateCreated       time.Time    `json:"date_created,omitempty"`
}

func (request *PaymentProcessRequest) Validate() *rest_errors.RestErr {
	if request.Id == "" {
		return rest_errors.NewBadRequestError("invalid id")
	}
	if len(request.UserSecrets) > 0 {
		for _, requestSecret := range request.UserSecrets {
			if requestSecret.Key != "" && strings.TrimSpace(requestSecret.Value) == "" {
				return rest_errors.NewBadRequestError(fmt.Sprintf("invalid %s", requestSecret.Human))
			}
		}
	}
	return nil
}

type PaymentResult struct {
	Id                 primitive.ObjectID                   `json:"id,omitempty" bson:"_id, omitempty"`
	Payer              customer_dto.Customer                `json:"payer"`
	Payee              customer_dto.Customer                `json:"payee"`
	UserSecrets        []UserSecret                         `json:"user_secrets,omitempty"`
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
	DateCreated        time.Time                            `json:"date_created,omitempty"`
	LastUpdated        time.Time                            `json:"last_updated,omitempty"`
}

type WalletPaymentRequest struct {
	PayerWallet wallet_dao.Wallet    `json:"payer_wallet"`
	PayeeWallet wallet_dao.Wallet    `json:"payee_wallet"`
	Payment     PublicPaymentRequest `json:"payment"`
}

func (request *WalletPaymentRequest) Validate() *rest_errors.RestErr {
	if payerErr := request.PayerWallet.Validate(); payerErr != nil {
		return rest_errors.NewBadRequestError("invalid payer data")
	}
	if payeeErr := request.PayeeWallet.Validate(); payeeErr != nil {
		return rest_errors.NewBadRequestError("invalid payee data")
	}
	if paymentErr := request.Payment.Payee.Validate(); paymentErr != nil {
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
