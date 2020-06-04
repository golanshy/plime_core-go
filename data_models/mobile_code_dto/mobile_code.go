package mobile_code_dto

import (
	"github.com/golanshy/plime_core-go/utils/date_utils"
	"math/rand"
	"time"
)

const (
	expirationTime = 15
)

type MobileCode struct {
	Code        int    `json:"code"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
	Expires     int64  `json:"expires"`
}

func GetMobileCodeByEmail(email string) *MobileCode {
	return &MobileCode{
		Code:        0,
		Email:       email,
		DateCreated: date_utils.GetNowDBFormat(),
		Expires:     time.Now().UTC().Add(expirationTime * time.Minute).Unix(),
	}
}

func (eat *MobileCode) GenerateNumber() {
	rand.Seed(time.Now().UnixNano())
	min := 100000
	max := 999999
	eat.Code = rand.Intn(max-min+1) + min
}

func (eat *MobileCode) IsExpired() bool {
	return time.Unix(eat.Expires, 0).Before(time.Now().UTC())
}
