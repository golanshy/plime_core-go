package rest_errors

import (
	"encoding/json"
	"errors"
	"net/http"
)

type RestErr interface {
	Message() string
	Status() int
	Error() string
	Causes() []interface{}
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
	return apiError, nil
}

func NewBadRequestError(message string) RestErr {
	return &restErr{
		message: message,
		status:  http.StatusBadRequest,
		error:   "bad_request",
	}
}

func NewNotFoundError(message string) RestErr {
	return &restErr{
		message: message,
		status:  http.StatusNotFound,
		error:   "not_found",
	}
}

func NewInternalServerError(message string, err error) RestErr {
	return &restErr{
		message: message,
		status:  http.StatusInternalServerError,
		error:   "internal_server_error",
		causes:  []interface{}{err},
	}
}

func NewUnauthorizedError(message string) RestErr {
	return &restErr{
		message: message,
		status:  http.StatusUnauthorized,
		error:   "unauthorized_access",
	}
}

func NewForbiddenEmailVerificationError(message string) RestErr {
	return &restErr{
		message: message,
		status:  http.StatusForbidden,
		error:   "forbidden_access",
	}
}
