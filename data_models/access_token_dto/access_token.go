package access_token_dto

import (
	"github.com/golanshy/plime_core-go/logger"
	"github.com/golanshy/plime_core-go/utils/date_utils"
	"github.com/golanshy/plime_core-go/utils/rest_errors"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	TokenTypeBearer            = "bearer"
	GrantTypeClientCredentials = "client_credentials"
	GrantTypeRefreshToken      = "refresh_token"
	GrantTypePassword          = "password"
)

var (
	expirationTime int64 = 24
)

func init() {
	expirationTimeFromEnv := os.Getenv("AUTH_ACCESS_TOKEN_EXPIRATION_TIME")

	var err error
	if expirationTimeFromEnv != "" {
		expirationTime, err = strconv.ParseInt(expirationTimeFromEnv, 10, 64)
		if err != nil {
			expirationTime = 24
		}
	}
}

type AccessTokenRequest struct {
	TokenType string `json:"token_type"`
	GrantType string `json:"grant_type"`
	Scope     string `json:"scope"`

	// User for password grant type
	Username string `json:"username"`
	Password string `json:"password"`

	// User for client_credentials_repo grant type
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`

	// User for refresh_token grant type
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func (request *AccessTokenRequest) Validate() *rest_errors.RestErr {
	switch request.GrantType {
	case GrantTypeClientCredentials:
		if request.ClientId == "" {
			logger.Error("error when trying to validate client_credentials_repo invalid client_id", nil)
			return rest_errors.NewBadRequestError("invalid client_id")
		}
		if request.ClientSecret == "" {
			logger.Error("error when trying to validate client_credentials_repo invalid client_secret", nil)
			return rest_errors.NewBadRequestError("invalid client_secret")
		}
		break
	case GrantTypeRefreshToken:
		if request.AccessToken == "" {
			logger.Error("error when trying to validate client_credentials_repo invalid access_token", nil)
			return rest_errors.NewBadRequestError("invalid access token")
		}
		if request.RefreshToken == "" {
			logger.Error("error when trying to validate client_credentials_repo invalid refresh_token", nil)
			return rest_errors.NewBadRequestError("invalid refresh token")
		}
		break
	case GrantTypePassword:
		if request.Username == "" {
			logger.Error("error when trying to validate client_credentials_repo invalid username", nil)
			return rest_errors.NewBadRequestError("invalid username")
		}
		if request.Password == "" {
			logger.Error("error when trying to validate client_credentials_repo invalid password", nil)
			return rest_errors.NewBadRequestError("invalid password")
		}
	default:
		logger.Error("error when trying to validate client_credentials_repo", nil)
		return rest_errors.NewBadRequestError("invalid grant type")
	}
	return nil
}

type AccessToken struct {
	Token       string `json:"token"`
	DateCreated string `json:"date_created"`
	Expires     int64  `json:"expires"`
	DateRevoked string `json:"date_revoked,omitempty"`
}

// Web Frontend ClientId: 123
// Android application ClientId: 234

func GetNewAccessToken() *AccessToken {
	return &AccessToken{
		DateCreated: date_utils.GetNowString(),
		Expires:     time.Now().UTC().Add(time.Duration(expirationTime * int64(time.Hour))).Unix(),
	}
}

func (at AccessToken) IsExpired() bool {
	return time.Unix(at.Expires, 0).Before(time.Now().UTC())
}

func (at *AccessToken) Validate() *rest_errors.RestErr {
	at.Token = strings.TrimSpace(at.Token)
	if at.Token == "" {
		logger.Error("error when trying to validate access_token invalid access token id", nil)
		return rest_errors.NewBadRequestError("invalid access token id")
	}
	if at.Expires <= 0 {
		logger.Error("error when trying to validate access_token invalid expiration time", nil)
		return rest_errors.NewBadRequestError("invalid expiration time")
	}
	return nil
}

func (at *AccessToken) UpdateExpirationTime() {
	at.Expires = time.Now().UTC().Add(time.Duration(expirationTime * int64(time.Hour))).Unix()
}
