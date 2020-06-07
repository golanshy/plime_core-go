package rails_bank_transaction_dto

import "time"

type RailsBankInterLedgerTransactionRequest struct {
	Amount        float64 `json:"amount"`
	PayerWalletID string  `json:"ledger_from_id"`
	PayeeWalletId string  `json:"ledger_to_id"`
}

type RailsBankInterLedgerTransactionResponse struct {
	TransactionId string `json:"transaction_id"`
}

type RailsBankTransactionResponse struct {
	TransactionId            string                       `json:"transaction_id"`
	TransactionStatus        string                       `json:"transaction_status"`
	TransactionType          string                       `json:"transaction_type"`
	CreatedAt                *time.Time                   `json:"created_at"`
	PartnerProduct           string                       `json:"partner_product"`
	TransactionFee           float64                      `json:"transaction_fee"`
	AssetClass               string                       `json:"asset_class"`
	AssetType                string                       `json:"asset_type"`
	FailureReasons           []string                     `json:"failure_reasons"`
	AmountBeneficiaryAccount float64                      `json:"amount_beneficiary_account"`
	Amount                   float64                      `json:"amount"`
	AmountLedgerFrom         float64                      `json:"amount_ledger_from"`
	AmountLocalCurrency      float64                      `json:"amount_local_currency"`
	SettlementDate           string                       `json:"settlement_date"`
	LedgerFromId             string                       `json:"ledger_from_id"`
	LedgerToId               string                       `json:"ledger_to_id"`
	PaymentMethod            string                       `json:"payment_method"`
	PaymentType              string                       `json:"payment_type"`
	Reference                string                       `json:"reference"`
	RejectionReasons         []string                     `json:"rejection_reasons"`
	SwiftChargeBearer        string                       `json:"swift_charge_bearer"`
	SwiftServiceLevel        string                       `json:"swift_service_level"`
	TransactionAuditNumber   string                       `json:"transaction_audit_number"`
	TransactionInfo          RailsBankTransactionInfo     `json:"transaction_info"`
	TransactionPrintout      RailsBankTransactionPrintout `json:"transaction_printout"`
}

type RailsBankTransactionInfo struct {
	Amount                    float64 `json:"amount"`
	ChargeBearer              string  `json:"charge_bearer"`
	ClearingSystem            string  `json:"clearing_system"`
	ClearingSystemProprietary string  `json:"clearing_system_proprietary"`
	Currency                  string  `json:"currency"`
	EndToEndId                string  `json:"end_to_end_id"`
}

type RailsBankTransactionPrintout struct {
	BeneficiaryIban                       float64 `json:"beneficiaryiban"`
	PspOfSenderName                       string  `json:"pspofsendername"`
	PspAccountLocation                    string  `json:"pspaccountlocation"`
	PspAccountTandCsCountryOfJurisdiction string  `json:"pspaccounttandcscountryofjurisdiction"`
	UltimateSenderName                    string  `json:"ultimatesendername"`
	PaymentOnBehalfOfType                 string  `json:"paymentonbehalfoftype"`
	PspOfUltimateBeneName                 string  `json:"pspofultimatebenename"`
	UltimateSenderAccountNumber           string  `json:"ultimatesenderaccountnumber"`
	BeneficiaryName                       string  `json:"beneficiaryname"`
	PaymentPartyType                      string  `json:"paymentpartytype"`
}

type TransactionsResultsResponse struct {
	Start   int64                           `json:"start"`
	Limit   int64                           `json:"limit"`
	Hits    int64                           `json:"hits"`
	Results *[]RailsBankTransactionResponse `json:"results,omitempty"`
}
