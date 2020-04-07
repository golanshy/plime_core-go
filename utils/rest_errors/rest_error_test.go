package rest_errors

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestNewInternalServerError(t *testing.T) {
	err := NewInternalServerError("this is the message", errors.New("this is the error"))
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "this is the message", err.Message)
	assert.EqualValues(t, "internal_server_error", err.Error)

	assert.NotNil(t, err.Causes)
	assert.EqualValues(t, 1, len(err.Causes))
	//assert.EqualValues(t, "this is the error", err.Causes()[0])
}

func TestNewBadRequestError(t *testing.T) {
	// TODO: Test!!!
}

func TestNewError(t *testing.T) {
	// TODO: Test!!!
}

func TestNewForbiddenEmailVerificationError(t *testing.T) {
	// TODO: Test!!!
}

func TestNewNotFoundError(t *testing.T) {
	// TODO: Test!!!
}

func TestNewRestError(t *testing.T) {
	// TODO: Test!!!
}

func TestNewRestErrorFromBytes(t *testing.T) {
	// TODO: Test!!!
}

func TestNewUnauthorizedError(t *testing.T) {
	// TODO: Test!!!
}

func TestRestErr_Causes(t *testing.T) {
	// TODO: Test!!!
}

func TestRestErr_Error(t *testing.T) {
	// TODO: Test!!!
}

func TestRestErr_Message(t *testing.T) {
	// TODO: Test!!!
}

func TestRestErr_Status(t *testing.T) {
	// TODO: Test!!!
}
