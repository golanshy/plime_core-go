package rails_bank_enduser_dto

import "github.com/golanshy/plime_core-go/data_models/rails_bank_models/rails_bank_ledger_dto"

type RailsBankEndUserRequest struct {
	Person  *RailsBankPerson  `json:"person,omitempty"`
	Company *RailsBankCompany `json:"company,omitempty"`
}

type RailsBankEndUserResponse struct {
	EndUserId                string                                    `json:"enduser_id"`
	EndUserStatus            string                                    `json:"enduser_status,omitempty"`
	Ledgers                  []rails_bank_ledger_dto.RailsBankLedgerId `json:"ledgers,omitempty"`
	Person                   *RailsBankPerson                          `json:"person,omitempty"`
	Company                  *RailsBankCompany                         `json:"company,omitempty"`
	CreatedAt                string                                    `json:"created_at,omitempty"`
	LastModifiedAt           string                                    `json:"last_modified_at,omitempty"`
	ScreeningMonitoredSearch bool                                      `json:"screening_monitored_search,omitempty"`
}

type RailsBankPerson struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type RailsBankCompany struct {
	Name                string                      `json:"name,omitempty"`
	LegalEntryNumber    string                      `json:"legal_entry_number,omitempty"`
	RegistrationNumber  string                      `json:"registration_number,omitempty"`
	Email               string                      `json:"email,omitempty"`
	Telephone           string                      `json:"telephone,omitempty"`
	Website             string                      `json:"web_site,omitempty"`
	RegistrationAddress *RailsBankRegisteredAddress `json:"registration_address,omitempty"`
}

type RailsBankRegisteredAddress struct {
	AddressNumber     string `json:"address_number,omitempty"`      // The house number
	AddressStreet     string `json:"address_street,omitempty"`      // The street name
	AddressRefinement string `json:"address_refinement,omitempty"`  // The address refinement, may include flat number, house name etc
	AddressCity       string `json:"address_city,omitempty"`        // The city or locality
	AddressRegion     string `json:"address_region,omitempty"`      // The state, province, prefecture, or region
	AddressIsoCountry string `json:"address_iso_country,omitempty"` // The ISO country code
	AddressPostalCode string `json:"address_postal_code,omitempty"` // The zip code or postal code
}
