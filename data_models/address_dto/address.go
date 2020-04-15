package address_dto

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