package payment_dto

import (
	"fmt"
	"github.com/golanshy/plime_core-go/data_models/transaction_dto"
	"github.com/golanshy/plime_core-go/data_models/user_dto"
	"github.com/golanshy/plime_core-go/utils/date_utils"
	"github.com/golanshy/plime_core-go/utils/rest_errors"
	"strings"
	"time"
)

type PaymentsRequest struct {
	Id        string    `json:"id"`
	Reference string    `json:"reference"`
	Details   string    `json:"details"`
	Payments  []Payment `json:"payments"`
	WebHook   WebHook   `json:"web_hook"`
}

type Payment struct {
	Id           string        `json:"id"`
	User         user_dto.User `json:"user"`
	UserSecrets  []UserSecret  `json:"user_secrets"`
	Reference    string        `json:"reference"`
	Details      string        `json:"details"`
	Amount       float64       `json:"amount"`
	CurrencyCode string        `json:"currency_code"`
	SendOn       string        `json:"send_on"`
	ArriveBy     string        `json:"arrive_by"`
}

func (payment *Payment) Validate() *rest_errors.RestErr {
	if userErr := payment.User.Validate(); userErr != nil {
		return rest_errors.NewBadRequestError("invalid user data")
	}
	if payment.Amount <= 0 {
		return rest_errors.NewBadRequestError("invalid amount")
	}
	if strings.TrimSpace(payment.CurrencyCode) == "" {
		return rest_errors.NewBadRequestError("invalid currency code")
	}
	if len(payment.UserSecrets) > 0 {
		for _, secret := range payment.UserSecrets {
			if strings.TrimSpace(secret.Key) == "" {
				return rest_errors.NewBadRequestError("invalid user secret key")
			}
			if strings.TrimSpace(secret.Human) == "" {
				return rest_errors.NewBadRequestError("invalid user secret human")
			}
			if strings.TrimSpace(secret.Value) == "" {
				return rest_errors.NewBadRequestError("invalid user secret value")
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
		payment.User.Email = strings.TrimSpace(payment.User.Email)
		if payment.User.Email == "" {
			return rest_errors.NewBadRequestError(fmt.Sprintf("invalid email address for payment %d", index))
		}
		if payment.Reference == "" {
			return rest_errors.NewBadRequestError(fmt.Sprintf("invalid reference data for payment %d", index))
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
}

func (request *PaymentsResponse) Validate() *rest_errors.RestErr {
	//if request.Email == "" {
	//	return rest_errors.NewBadRequestError("invalid email address")
	//}
	return nil
}
