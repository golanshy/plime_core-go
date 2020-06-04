package transaction_dto

import "github.com/golanshy/plime_core-go/data_models/rails_bank_models/rails_bank_transaction_dto"

type Transaction struct {
}

type TransactionResult struct {
	Id                  string                                                  `json:"id"`
	Status              string                                                  `json:"status"`
	CreatedAt           string                                                  `json:"created_at"`
	TransactionFee      float64                                                 `json:"transaction_fee,omitempty"`
	CurrencyCode        string                                                  `json:"currency_code"`
	FailureReasons      []string                                                `json:"failure_reasons,omitempty"`
	Amount              float64                                                 `json:"amount"`
	AmountLocalCurrency float64                                                 `json:"amount_local_currency,omitempty"`
	PayerWalletId       string                                                  `json:"payer_wallet_id"`
	PayeeWalletId       string                                                  `json:"payee_wallet_id"`
	Details             rails_bank_transaction_dto.RailsBankTransactionResponse `json:"details,omitempty"`
}

const (
	StatusPending  string = "pending"
	StatusFailed   string = "failed"
	StatusSuccess  string = "success"
	StatusDeclined string = "declined"
)
