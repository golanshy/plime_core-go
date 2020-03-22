package rest_errors

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type RestErr interface {
	Message() string
	Status() int
	Error() string
	Causes() []interface{}
}

type restErr struct {
	message string        `json:"message"`
	status  int           `json:"status"`
	error   string        `json:"error"`
	causes  []interface{} `json:"causes"`
}

func (e restErr) restErr() string {
	return fmt.Sprintf("message: %s - status %d - error: %s - causes: [ %v]", e.message, e.status, e.error, e.causes)
}

func (e restErr) Message() string {
	return e.message
}
func (e restErr) Status() int {
	return e.status
}
func (e restErr) Error() string {
	return e.error
}
func (e restErr) Causes() []interface{} {
	return e.causes
}

func NewError(message string) error {
	return errors.New(message)
}

func NewRestError(message string, status int, error string, causes []interface{}) RestErr {
	return &restErr{
		message: message,
		status:  status,
		error:   error,
		causes:  causes,
	}
}

func NewRestErrorFromBytes(bytes []byte) (RestErr, error) {
	var apiError restErr
	if err := json.Unmarshal(bytes, &apiError); err != nil {
		return nil, NewError("invalid json")
	}
	return apiError, nil
}

func NewBadRequestError(message string) RestErr {
	return restErr{
		message: message,
		status:  http.StatusBadRequest,
		error:   "bad_request",
	}
}

func NewNotFoundError(message string) RestErr {
	return restErr{
		message: message,
		status:  http.StatusNotFound,
		error:   "not_found",
	}
}

func NewInternalServerError(message string, err error) RestErr {
	return restErr{
		message: message,
		status:  http.StatusInternalServerError,
		error:   "internal_server_error",
		causes:  []interface{}{err},
	}
}

func NewUnauthorizedError(message string) RestErr {
	return restErr{
		message: message,
		status:  http.StatusUnauthorized,
		error:   "unauthorized_access",
	}
}

func NewForbiddenEmailVerificationError(message string) RestErr {
	return restErr{
		message: message,
		status:  http.StatusForbidden,
		error:   "forbidden_access",
	}
}
