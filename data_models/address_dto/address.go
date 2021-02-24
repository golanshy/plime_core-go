package address_dto

import (
	"github.com/golanshy/plime_core-go/utils/rest_errors"
	"strings"
)

type Address struct {
	FirstLine  string  `json:"first_line,omitempty"`
	SecondLine string  `json:"second_line,omitempty"`
	ThirdLine  string  `json:"third_line,omitempty"`
	Town       string  `json:"town,omitempty"`
	PostCode   string  `json:"post_code,omitempty"`
	County     string  `json:"county,omitempty"`
	State      string  `json:"state,omitempty"`
	Country    string  `json:"country,omitempty"`
	Latitude   float64 `json:"latitude"`
	Longitude  float64 `json:"longitude"`
}

func (address *Address) Trim() {
	address.FirstLine = strings.TrimSpace(address.FirstLine)
	address.SecondLine = strings.TrimSpace(address.SecondLine)
	address.ThirdLine = strings.TrimSpace(address.ThirdLine)
	address.Town = strings.TrimSpace(address.Town)
	address.PostCode = strings.TrimSpace(address.PostCode)
	address.County = strings.TrimSpace(address.County)
	address.State = strings.TrimSpace(address.State)
	address.Country = strings.TrimSpace(address.Country)
}

func (address *Address) Validate() *rest_errors.RestErr {
	if address.FirstLine == "" {
		return rest_errors.NewBadRequestError("invalid address first line field")
	}
	if address.Town == "" {
		return rest_errors.NewBadRequestError("invalid address town field")
	}
	if address.PostCode == "" {
		return rest_errors.NewBadRequestError("invalid address postcode field")
	}
	if address.Country == "" {
		return rest_errors.NewBadRequestError("invalid address country field")
	}
	return nil
}
