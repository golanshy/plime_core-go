package access_token_dto

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestAccessTokenConstants(t *testing.T) {
	assert.EqualValues(t, 24, expirationTime)
	assert.EqualValues(t, "client_credentials_repo", GrantTypeClientCredentials)
	assert.EqualValues(t, "refresh_token", GrantTypeRefreshToken)
	assert.EqualValues(t, "bearer", TokenTypeBearer)
}

func TestGetNewAccessTokenByUserId(t *testing.T) {
	at := GetNewAccessTokenByUserId("123")
	assert.EqualValues(t, "", at.AccessToken, "new access token should not have defined access token id")
	assert.EqualValues(t, 123, at.UserId, "new access token should not have an associated user id")
	assert.True(t, at.UserId == "123", "new access token should not have an associated user id")
	assert.False(t, at.IsExpired(), "brand new access token should not be expired")
}

func TestGetNewAccessTokenByClientId(t *testing.T) {
	at := GetNewAccessTokenByClientId("123")
	assert.EqualValues(t, "", at.AccessToken, "new access token should not have defined access token id")
	assert.EqualValues(t, "123", at.ClientId, "new access token should not have an associated user id")
	assert.True(t, at.UserId == "", "new access token should not have an associated user id")
	assert.False(t, at.IsExpired(), "brand new access token should not be expired")
}

func TestAccessTokenIsExpired(t *testing.T) {
	at := AccessToken{}
	assert.True(t, at.IsExpired(), "empty access token should be expired by default as Expires = 0")

	at.Expires = time.Now().UTC().Add(3 * time.Hour).Unix()
	assert.False(t, at.IsExpired(), "should expire in 3 hours")
}
