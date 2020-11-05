package client_credentials_dto

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestConstantsGrantTypeClientCredentials(t *testing.T) {
	assert.EqualValues(t, "client_credentials", GrantTypeClientCredentials)
}

func TestConstantsSecretLength(t *testing.T) {
	assert.EqualValues(t, 64, SecretLength)
}

func TestClientCredentialsCreateRequest_ValidateClientNameFailed(t *testing.T) {
	cr := ClientCredentialsCreateRequest{
		ClientName:   "   ",
		AppName:      "abc",
		AppDetails:   "def",
	}

	err := cr.Validate()
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusBadRequest, err.Status)
	assert.EqualValues(t, "invalid client_name", err.Message)
	assert.EqualValues(t, "bad_request", err.Error)
}

func TestClientCredentialsCreateRequest_ValidateAppNameFailed(t *testing.T) {
	cr := ClientCredentialsCreateRequest{
		ClientName:   "456",
		AppName:      "   ",
		AppDetails:   "def",
	}

	err := cr.Validate()
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusBadRequest, err.Status)
	assert.EqualValues(t, "invalid app_name", err.Message)
	assert.EqualValues(t, "bad_request", err.Error)
}

func TestClientCredentialsCreateRequest_ValidateAppDetailsFailed(t *testing.T) {
	cr := ClientCredentialsCreateRequest{
		ClientName:   "456",
		AppName:      "abc",
		AppDetails:   "    ",
	}

	err := cr.Validate()
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusBadRequest, err.Status)
	assert.EqualValues(t, "invalid app_details", err.Message)
	assert.EqualValues(t, "bad_request", err.Error)
}

func TestClientCredentialsCreateRequest_ValidateDateCreatedSuccess(t *testing.T) {
	cr := ClientCredentialsCreateRequest{
		ClientName:   "   456",
		AppName:      "   abc",
		AppDetails:   "   def",
	}

	err := cr.Validate()
	assert.Nil(t, err)
}

func TestClientCredentials_ValidateClientIdFailed(t *testing.T) {
	cr := ClientCredentials{
		GrantType:    GrantTypeClientCredentials,
		ClientId:     "  ",
		ClientName:   "456",
		AppName:      "abc",
		AppDetails:   "def",
		ClientSecret: "123",
		DateCreated:  "1",
	}

	err := cr.Validate()
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusBadRequest, err.Status)
	assert.EqualValues(t, "invalid client_id", err.Message)
	assert.EqualValues(t, "bad_request", err.Error)
}

func TestClientCredentials_ValidateClientNameFailed(t *testing.T) {
	cr := ClientCredentials{
		GrantType:    GrantTypeClientCredentials,
		ClientId:     "123",
		ClientName:   "  ",
		AppName:      "abc",
		AppDetails:   "def",
		ClientSecret: "123",
		DateCreated:  "1",
	}

	err := cr.Validate()
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusBadRequest, err.Status)
	assert.EqualValues(t, "invalid client_name", err.Message)
	assert.EqualValues(t, "bad_request", err.Error)
}

func TestClientCredentials_ValidateAppNameFailed(t *testing.T) {
	cr := ClientCredentials{
		GrantType:    GrantTypeClientCredentials,
		ClientId:     "123",
		ClientName:   "456",
		AppName:      "   ",
		AppDetails:   "def",
		ClientSecret: "123",
		DateCreated:  "1",
	}

	err := cr.Validate()
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusBadRequest, err.Status)
	assert.EqualValues(t, "invalid app_name", err.Message)
	assert.EqualValues(t, "bad_request", err.Error)
}

func TestClientCredentials_ValidateAppDetailsFailed(t *testing.T) {
	cr := ClientCredentials{
		GrantType:    GrantTypeClientCredentials,
		ClientId:     "123",
		ClientName:   "456",
		AppName:      "abc",
		AppDetails:   "  ",
		ClientSecret: "123",
		DateCreated:  "1",
	}

	err := cr.Validate()
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusBadRequest, err.Status)
	assert.EqualValues(t, "invalid app_details", err.Message)
	assert.EqualValues(t, "bad_request", err.Error)
}

func TestClientCredentials_ValidateClientSecretFailed(t *testing.T) {
	cr := ClientCredentials{
		GrantType:    GrantTypeClientCredentials,
		ClientId:     "123",
		ClientName:   "456",
		AppName:      "abc",
		AppDetails:   "def",
		ClientSecret: "   ",
		DateCreated:  "1",
	}

	err := cr.Validate()
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusBadRequest, err.Status)
	assert.EqualValues(t, "invalid client_secret", err.Message)
	assert.EqualValues(t, "bad_request", err.Error)
}

func TestClientCredentials_ValidateGrantTypeFailed(t *testing.T) {
	cr := ClientCredentials{
		GrantType:    "    ",
		ClientId:     "123",
		ClientName:   "456",
		AppName:      "abc",
		AppDetails:   "def",
		ClientSecret: "123",
		DateCreated:  "1",
	}

	err := cr.Validate()
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusBadRequest, err.Status)
	assert.EqualValues(t, "invalid grant_type", err.Message)
	assert.EqualValues(t, "bad_request", err.Error)
}

func TestClientCredentials_ValidateDateCreatedFailed(t *testing.T) {
	cr := ClientCredentials{
		GrantType:    GrantTypeClientCredentials,
		ClientId:     "123",
		ClientName:   "456",
		AppName:      "abc",
		AppDetails:   "def",
		ClientSecret: "123",
		DateCreated:  "1",
	}

	err := cr.Validate()
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusBadRequest, err.Status)
	assert.EqualValues(t, "invalid date_created", err.Message)
	assert.EqualValues(t, "bad_request", err.Error)
}

func TestClientCredentials_ValidateDateCreatedSuccess(t *testing.T) {
	cr := ClientCredentials{
		GrantType:    GrantTypeClientCredentials,
		ClientId:     "123",
		ClientName:   "456",
		AppName:      "abc",
		AppDetails:   "def",
		ClientSecret: "123",
		DateCreated:  "1",
	}

	err := cr.Validate()
	assert.Nil(t, err)
}

func TestClientCredentialsRequest_Validate(t *testing.T) {

}

func TestGetNewClientCredentialsByClientId(t *testing.T) {

}

func TestGenerateSecret(t *testing.T) {

}
