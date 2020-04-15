package payment_dto

import (
	"fmt"
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
}

type Payment struct {
	Id           string  `json:"id"`
	Email        string  `json:"email"`
	Reference    string  `json:"reference"`
	Details      string  `json:"details"`
	Amount       float64 `json:"amount"`
	CurrencyCode string  `json:"currency_code"`
	SendOn       string  `json:"send_on"`
	ArriveBy     string  `json:"arrive_by"`
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
		payment.Email = strings.TrimSpace(payment.Email)
		payment.Reference = strings.TrimSpace(payment.Reference)
		payment.Details = strings.TrimSpace(payment.Details)
		payment.ArriveBy = strings.TrimSpace(payment.ArriveBy)
		payment.SendOn = strings.TrimSpace(payment.SendOn)
		payment.CurrencyCode = strings.TrimSpace(payment.CurrencyCode)
		if payment.Email == "" {
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
	Id              string          `json:"id"`
	Reference       string          `json:"reference"`
	Details         string          `json:"details"`
	PaymentsResults []PaymentResult `json:"payments_results"`
	FailureDetails string               `json:"failure_details"`
	Error          *rest_errors.RestErr `json:"error"`
}

type PaymentResult struct {
	Id             string               `json:"id"`
	Email          string               `json:"email"`
	Reference      string               `json:"reference"`
	Details        string               `json:"details"`
	Amount         float64              `json:"amount"`
	CurrencyCode   string               `json:"currency_code"`
	SendOn         string               `json:"send_on"`
	ArriveBy       string               `json:"arrive_by"`
	Status         string               `json:"status"`
	FailureDetails string               `json:"failure_details"`
	Error          *rest_errors.RestErr `json:"error"`
}

func (request *PaymentsResponse) Validate() *rest_errors.RestErr {
	//if request.Email == "" {
	//	return rest_errors.NewBadRequestError("invalid email address")
	//}
	return nil
}