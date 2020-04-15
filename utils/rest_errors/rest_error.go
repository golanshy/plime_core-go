package rest_errors

import (
	"encoding/json"
	"errors"
	"net/http"
)

type RestErr struct {
	Message string        `json:"message"`
	Status  int           `json:"status"`
	Error   string        `json:"error"`
	Causes  []interface{} `json:"causes"`
}

func NewRestErrorFromBytes(bytes []byte) (*RestErr, error) {
	var apiError RestErr
	if err := json.Unmarshal(bytes, &apiError); err != nil {
		return nil, errors.New("invalid json")
	}
	return &apiError, nil
}

func NewRestError(message string, status int, error string) *RestErr {
	return &RestErr{
		Message: message,
		Status:  status,
		Error:   error,
	}
}

func NewBadRequestError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Status:  http.StatusBadRequest,
		Error:   "bad_request",
	}
}

func NewNotFoundError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Status:  http.StatusNotFound,
		Error:   "not_found",
	}
}

func NewInternalServerError(message string, err error) *RestErr {
	return &RestErr{
		Message: message,
		Status:  http.StatusInternalServerError,
		Error:   "internal_server_error",
		Causes:  []interface{}{err},
	}
}

func NewUnauthorizedError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Status:  http.StatusUnauthorized,
		Error:   "unauthorized_access",
	}
}

func NewForbiddenError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Status:  http.StatusForbidden,
		Error:   "forbidden_access",
	}
}

func NewConflictError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Status:  http.StatusConflict,
		Error:   "conflict_error",
	}
}



