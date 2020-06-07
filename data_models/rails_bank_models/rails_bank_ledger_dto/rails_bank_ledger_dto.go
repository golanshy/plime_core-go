package rails_bank_ledger_dto

import "time"

type RailsBankLedgerRequest struct {
	AssetClass            string   `json:"asset_class"`
	AssetType             string   `json:"asset_type"`
	HolderId              string   `json:"holder_id"`
	LedgerPrimaryUseTypes []string `json:"ledger_primary_use_types"`
	CountryCode           string   `json:"ledger_t_and_cs_country_of_jurisdiction"`
	LedgerType            string   `json:"ledger_type"`
	WhoOwns               string   `json:"ledger_who_owns_assets"`
	PartnerProduct        string   `json:"partner_product"`
}

type RailsBankLedgerResponse struct {
	LedgerId        string     `json:"ledger_id"`
	HolderId        string     `json:"holder_id"`
	PartnerRef      string     `json:"partner_ref"`
	PartnerProduct  string     `json:"partner_product"`
	CurrencyCode    string     `json:"asset_type"` // Iso 4217 https://en.wikipedia.org/wiki/ISO_4217
	Iban            string     `json:"iban"`
	BicSwift        string     `json:"bic_swift"`
	WalletStatus    string     `json:"ledger_status"`
	IbanStatus      string     `json:"ledger_iban_status"`
	Amount          float64    `json:"amount"`
	CreatedAt       *time.Time `json:"created_at"`
	CountryCode     string     `json:"ledger_t_and_cs_country_of_jurisdiction"`
	UkSortCode      string     `json:"uk_sort_code"`
	UkAccountNumber string     `json:"uk_account_number"`
}
