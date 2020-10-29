package kyb_dto

const (
	KybNotStarted                   string = "kyb_status_not_started"
	KybStatusInProgress             string = "kyb_status_in_progress"
	KybStatusInReview               string = "kyb_status_in_review"
	KybStatusAdditionalDataRequired string = "kyb_status_additional_data_required"
	KybStatusApproved               string = "kyb_status_approved"
	KybStatusDeclined               string = "kyb_status_declined"
	KybStatusBanned                 string = "kyb_status_banned"

	KybLevel0 int = 0
	KybLevel1 int = 1
	KybLevel2 int = 2
	KybLevel3 int = 3
	KybLevel4 int = 4
	KybLevel5 int = 5
)

type KybStatus struct {
	CustomerId             string        `json:"customer_id,omitempty"`
	Status                 string        `json:"status,omitempty"`
	LevelApproved          int           `json:"level_approved,omitempty"`
	AdditionalDataRequired []KybDocument `json:"additional_data_required,omitempty"`
}

func NewKybStatus(customerId string) KybStatus {
	return KybStatus{
		CustomerId:    customerId,
		Status:        KybNotStarted,
		LevelApproved: KybLevel0,
	}
}
