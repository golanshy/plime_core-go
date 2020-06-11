package rails_bank_enduser_dto

import "github.com/golanshy/plime_core-go/data_models/rails_bank_models/rails_bank_ledger_dto"

type RailsBankEndUserRequest struct {
	Person RailsBankPerson `json:"person"`
}

type RailsBankEndUserResponse struct {
	EndUserId                string                                    `json:"enduser_id"`
	EndUserStatus            string                                    `json:"enduser_status,omitempty"`
	Ledgers                  []rails_bank_ledger_dto.RailsBankLedgerId `json:"ledgers,omitempty"`
	Person                   RailsBankPerson                           `json:"person,omitempty"`
	CreatedAt                string                                    `json:"created_at,omitempty"`
	LastModifiedAt           string                                    `json:"last_modified_at,omitempty"`
	ScreeningMonitoredSearch bool                                      `json:"screening_monitored_search,omitempty"`
}

type RailsBankPerson struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}
