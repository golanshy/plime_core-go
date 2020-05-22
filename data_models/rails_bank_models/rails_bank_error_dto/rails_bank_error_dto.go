package rails_bank_error_dto

import (
	"encoding/json"
	"github.com/golanshy/plime_core-go/utils/rest_errors"
)

type RailsBankError struct {
	Error string `json:"error"`
	ErrorDetails string `json:"error_details"`
}

func NewRestErrorFromBytes(bytes []byte) (*rest_errors.RestErr, error) {

	var apiError error_dto.RailsBankError
	if err := json.Unmarshal(bytes, &apiError); err != nil {
		return nil, errors.New("invalid json")
	}
	var restError = &rest_errors.RestErr{
		Message: apiError.Error,
		Status:  0,
		Error:   apiError.Error,
		Causes:  nil,
	}
	return restError, nil
}