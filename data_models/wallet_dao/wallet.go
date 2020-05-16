package wallet_dao

type Wallet struct {
	Id           string `json:"id"`
	Type         int    `json:"type"`
	CurrencyCode string `json:"currency_code"` // Iso 4217 https://en.wikipedia.org/wiki/ISO_4217
}

// WalletType:
// 0 = ???
