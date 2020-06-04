package address_dto

import "strings"

type Address struct {
	FirstLine  string `json:"first_line,omitempty"`
	SecondLine string `json:"second_line,omitempty"`
	ThirdLine  string `json:"third_line,omitempty"`
	Town       string `json:"town,omitempty"`
	PostCode   string `json:"post_code,omitempty"`
	County     string `json:"county,omitempty"`
	State      string `json:"state,omitempty"`
	Country    string `json:"country,omitempty"`
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
