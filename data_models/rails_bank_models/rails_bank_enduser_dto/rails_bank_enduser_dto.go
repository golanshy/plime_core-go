package rails_bank_enduser_dto

type RailsBankEndUserRequest struct {
	Person RailsBankPerson `json:"person"`
}

type RailsBankEndUserResponse struct {
	EndUserId string `json:"enduser_id"`
}

type RailsBankPerson struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}
