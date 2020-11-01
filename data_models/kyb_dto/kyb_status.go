package kyb_dto

const (
	KybNotStarted                   string = "kyb_status_not_started"
	KybStatusInProgress             string = "kyb_status_in_progress"
	KybStatusInReview               string = "kyb_status_in_review"
	KybStatusAdditionalDataRequired string = "kyb_status_additional_data_required"
	KybStatusApproved               string = "kyb_status_approved"
	KybStatusDeclined               string = "kyb_status_declined"
	KybStatusBanned                 string = "kyb_status_banned"

	KybLevel0 int64 = 0
	KybLevel1 int64 = 1
	KybLevel2 int64 = 2
	KybLevel3 int64 = 3
	KybLevel4 int64 = 4
	KybLevel5 int64 = 5
)

type KybStatus struct {
	CustomerId             string        `json:"customer_id"`
	Status                 string        `json:"status"`
	LevelApproved          int64          `json:"level_approved"`
	AdditionalDataRequired []KybDocument `json:"additional_data_required,omitempty"`
}

func NewKybStatus(customerId string) KybStatus {
	return KybStatus{
		CustomerId:    customerId,
		Status:        KybNotStarted,
		LevelApproved: KybLevel0,
	}
}
