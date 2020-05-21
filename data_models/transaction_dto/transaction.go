package transaction_dto

type Transaction struct {
}

type TransactionResult struct {
	Status string `json:"status"`
	Error  string `json:"error"`
}

const (
	TransactionStatusPending string = "pending"
	TransactionStatusFailed  string = "failed"
	TransactionStatusSuccess string = "success"
)
