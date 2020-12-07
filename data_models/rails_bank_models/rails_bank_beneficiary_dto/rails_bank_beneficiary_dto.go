package rails_bank_beneficiary_dto

import (
	"github.com/golanshy/plime_core-go/data_models/rails_bank_models/rails_bank_enduser_dto"
)

type RailsBankBeneficiaryRequest struct {
	HolderId        string                                   `json:"holder_id"`
	AssetClass      string                                   `json:"asset_class"`
	AssetType       string                                   `json:"asset_type"`
	UkAccountNumber string                                   `json:"uk_account_number,omitempty"`
	UkSortCode      string                                   `json:"uk_sort_code,omitempty"`
	Iban            string                                   `json:"iban,omitempty"`
	BicSwift        string                                   `json:"bic_swift,omitempty"`
	BankCountry     string                                   `json:"bank_country,omitempty"`
	BankName        string                                   `json:"bank_name,omitempty"`
	Person          *rails_bank_enduser_dto.RailsBankPerson  `json:"person,omitempty"`
	Company         *rails_bank_enduser_dto.RailsBankCompany `json:"company,omitempty"`
	BeneficiaryMeta *RailsBankBeneficiaryMeta                `json:"beneficiary_meta,omitempty"`
}

type RailsBankBeneficiaryId struct {
	BeneficiaryId string `json:"beneficiary_id"`
}

type RailsBankBeneficiaryMeta struct {
	Details string `json:"details,omitempty"`
}

type RailsBankBeneficiaryResponse struct {
	HolderId          string                                   `json:"holder_id"`
	BeneficiaryId     string                                   `json:"beneficiary_id"`
	BeneficiaryStatus string                                   `json:"beneficiary_status"`
	AssetClass        string                                   `json:"asset_class"`
	AssetType         string                                   `json:"asset_type"`
	UkAccountNumber   string                                   `json:"uk_account_number,omitempty"`
	UkSortCode        string                                   `json:"uk_sort_code,omitempty"`
	Iban              string                                   `json:"iban,omitempty"`
	BicSwift          string                                   `json:"bic_swift,omitempty"`
	BankCountry       string                                   `json:"bank_country,omitempty"`
	BankName          string                                   `json:"bank_name,omitempty"`
	Person            *rails_bank_enduser_dto.RailsBankPerson  `json:"person,omitempty"`
	Company           *rails_bank_enduser_dto.RailsBankCompany `json:"company,omitempty"`
}
