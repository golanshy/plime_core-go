package email_auth_token_dto

import (
	"crypto/rand"
	"fmt"
	"github.com/golanshy/plime_core-go/utils/date_utils"
	"time"
)

const (
	expirationTime = 15
)

type AuthToken struct {
	EmailAuthToken string `json:"email_auth_token"`
	Email          string `json:"email"`
	DateCreated    string `json:"date_created"`
	Expires        int64  `json:"expires"`
}

func GetNewEmailAuthTokenByEmail(email string) *AuthToken {
	return &AuthToken{
		EmailAuthToken: "",
		Email:          email,
		DateCreated:    date_utils.GetNowDBFormat(),
		Expires:        time.Now().UTC().Add(expirationTime * time.Minute).Unix(),
	}
}

func (eat *AuthToken) Generate() {
	data := make([]byte, 256)
	rand.Read(data)
	eat.EmailAuthToken = fmt.Sprintf("%x", data)
}

func (eat *AuthToken) IsExpired() bool {
	return time.Unix(eat.Expires, 0).Before(time.Now().UTC())
}