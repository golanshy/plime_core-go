package wallet_dao

type Wallet struct {
	Id             string  `json:"id"`
	HolderId       string  `json:"holder_id"`
	PartnerRef     string  `json:"partner_ref"`
	PartnerProduct string  `json:"partner_product"`
	CurrencyCode   string  `json:"currency_code"` // Iso 4217 https://en.wikipedia.org/wiki/ISO_4217
	Iban           string  `json:"iban"`
	BicSwift       string  `json:"bic_swift"`
	WalletStatus   string  `json:"status"`
	IbanStatus     string  `json:"iban_status"`
	Amount         float64 `json:"amount"`
	CreatedAt      string  `json:"created_at"`
	CountryCode    string  `json:"country_code"`
	UkSortCode     string  `json:"uk_sort_code"`
	UkAccountNumber string `json:"uk_account_number"`
}
