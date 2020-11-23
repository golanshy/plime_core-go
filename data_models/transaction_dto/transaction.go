package transaction_dto

import (
	"github.com/golanshy/plime_core-go/data_models/rails_bank_models/rails_bank_transaction_dto"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Transaction struct {
}

type TransactionResult struct {
	Id                  primitive.ObjectID                                       `json:"id,omitempty" bson:"_id, omitempty"`
	BankingId           string                                                   `json:"banking_id"`
	Status              string                                                   `json:"status"`
	CreatedAt           time.Time                                                `json:"created_at"`
	TransactionFee      float64                                                  `json:"transaction_fee,omitempty"`
	CurrencyCode        string                                                   `json:"currency_code"`
	FailureReasons      []string                                                 `json:"failure_reasons,omitempty"`
	Amount              float64                                                  `json:"amount"`
	AmountLocalCurrency float64                                                  `json:"amount_local_currency,omitempty"`
	PayerWalletId       string                                                   `json:"payer_wallet_id"`
	PayeeWalletId       string                                                   `json:"payee_wallet_id"`
	Details             *rails_bank_transaction_dto.RailsBankTransactionResponse `json:"details,omitempty"`
}

const (
	StatusAwaitingPayee string = "awaiting_payee"
	StatusFailed        string = "failed"
	StatusCompleted     string = "completed"
	StatusDeclined      string = "declined"
)
