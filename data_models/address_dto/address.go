package address_dto

import "strings"

type Address struct {
	FirstLine  string `json:"first_line"`
	SecondLine string `json:"second_line"`
	ThirdLine  string `json:"third_line"`
	Town       string `json:"town"`
	PostCode   string `json:"post_code"`
	County     string `json:"county"`
	State      string `json:"state"`
	Country    string `json:"country"`
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