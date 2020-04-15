package wallet_dao

type Wallet struct {
	WalletType int `json:"wallet_type"`
	WalletId   string `json:"wallet_id"`
}

// WalletType:
// 0 = IFX Payments
