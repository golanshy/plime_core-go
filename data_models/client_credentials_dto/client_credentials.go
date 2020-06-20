package client_credentials_dto

import (
	"fmt"
	"github.com/golanshy/plime_core-go/utils/crypto_utils"
	"github.com/golanshy/plime_core-go/utils/date_utils"
	"github.com/golanshy/plime_core-go/utils/rest_errors"
	"math/rand"
	"strings"
)

const (
	GrantTypeClientCredentials = "client_credentials"
	SecretLength               = 32
	ClientIdLength             = 16
)

type ClientCredentialsCreateRequest struct {
	ClientName string `json:"client_name"`
	AppName    string `json:"app_name"`
	AppDetails string `json:"app_details"`
}

func (request *ClientCredentialsCreateRequest) Validate() *rest_errors.RestErr {

	request.ClientName = strings.TrimSpace(request.ClientName)
	request.AppName = strings.TrimSpace(request.AppName)
	request.AppDetails = strings.TrimSpace(request.AppDetails)

	if request.ClientName == "" {
		return rest_errors.NewBadRequestError("invalid client_name")
	}
	if request.AppName == "" {
		return rest_errors.NewBadRequestError("invalid app_name")
	}
	if request.AppDetails == "" {
		return rest_errors.NewBadRequestError("invalid app_details")
	}
	return nil
}

func NewClientCredentials(request ClientCredentialsCreateRequest) ClientCredentials {
	return ClientCredentials{
		ClientId:     GenerateClientID(ClientIdLength),
		ClientName:   request.ClientName,
		AppName:      request.AppName,
		AppDetails:   request.AppDetails,
		ClientSecret: crypto_utils.GenerateSecret(SecretLength),
		GrantType:    GrantTypeClientCredentials,
		DateCreated:  date_utils.GetNowDBFormat(),
	}
}

func GenerateClientID(length int) string {
	data := make([]byte, length)
	_, _ = rand.Read(data)
	return fmt.Sprintf("%x", data)
}

type ClientCredentialsRequest struct {
	GrantType string `json:"grant_type"`

	// User for client_credentials grant type
	ClientId string `json:"client_id"`
}

func (request *ClientCredentialsRequest) Validate() *rest_errors.RestErr {
	if request.ClientId == "" {
		return rest_errors.NewBadRequestError("invalid client_id")
	}
	return nil
}

type ClientCredentials struct {
	GrantType string `json:"grant_type"`

	// User for client_credentials grant type
	ClientId     string `json:"client_id"`
	ClientName   string `json:"client_name"`
	AppName      string `json:"app_name"`
	AppDetails   string `json:"app_details"`
	ClientSecret string `json:"client_secret"`
	DateCreated  string  `json:"date_created"`
}

func (credentials *ClientCredentials) Validate() *rest_errors.RestErr {

	credentials.ClientId = strings.TrimSpace(credentials.ClientId)
	credentials.ClientName = strings.TrimSpace(credentials.ClientName)
	credentials.AppName = strings.TrimSpace(credentials.AppName)
	credentials.AppDetails = strings.TrimSpace(credentials.AppDetails)
	credentials.ClientSecret = strings.TrimSpace(credentials.ClientSecret)
	credentials.GrantType = strings.TrimSpace(credentials.GrantType)

	if credentials.ClientId == "" {
		return rest_errors.NewBadRequestError("invalid client_id")
	}
	if credentials.ClientName == "" {
		return rest_errors.NewBadRequestError("invalid client_name")
	}
	if credentials.AppName == "" {
		return rest_errors.NewBadRequestError("invalid app_name")
	}
	if credentials.AppDetails == "" {
		return rest_errors.NewBadRequestError("invalid app_details")
	}
	if credentials.ClientSecret == "" {
		return rest_errors.NewBadRequestError("invalid client_secret")
	}
	if credentials.GrantType == "" {
		return rest_errors.NewBadRequestError("invalid grant_type")
	}
	if credentials.DateCreated <= 0 {
		return rest_errors.NewBadRequestError("invalid date_created")
	}
	return nil
}
