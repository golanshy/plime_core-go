package access_token_dto

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestAccessTokenConstants(t *testing.T) {
	assert.EqualValues(t, int64(24), expirationTime)
	assert.EqualValues(t, "client_credentials", GrantTypeClientCredentials)
	assert.EqualValues(t, "refresh_token", GrantTypeRefreshToken)
	assert.EqualValues(t, "bearer", TokenTypeBearer)
}

func TestGetNewAccessTokenByUserId(t *testing.T) {
	at := GetNewAccessToken()
	assert.EqualValues(t, "", at.Token, "new access token should not have defined access token id")
	assert.False(t, at.IsExpired(), "brand new access token should not be expired")
}

func TestAccessTokenIsExpired(t *testing.T) {
	at := AccessToken{}
	assert.True(t, at.IsExpired(), "empty access token should be expired by default as Expires = 0")

	at.Expires = time.Now().UTC().Add(3 * time.Hour).Unix()
	assert.False(t, at.IsExpired(), "should expire in 3 hours")
}
