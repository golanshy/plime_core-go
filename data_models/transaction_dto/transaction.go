package transaction_dto

type Transaction struct {
}

type TransactionResult struct {
	Id                  string                       `json:"id"`
	Status              string                       `json:"status"`
	CreatedAt           string                       `json:"created_at"`
	TransactionFee      float64                      `json:"transaction_fee"`
	CurrencyCode        string                       `json:"currency_code"`
	FailureReasons      []string                     `json:"failure_reasons"`
	Amount              float64                      `json:"amount"`
	AmountLocalCurrency float64                      `json:"amount_local_currency"`
	PayerWalletId       string                       `json:"payer_wallet_id"`
	PayeeWalletId       string                       `json:"payee_wallet_id"`
	Details             RailsBankTransactionResponse `json:"details"`
}

const (
	StatusPending  string = "pending"
	StatusFailed   string = "failed"
	StatusSuccess  string = "success"
	StatusDeclined string = "declined"
)
