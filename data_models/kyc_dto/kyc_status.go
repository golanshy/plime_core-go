package kyc_dto

const (
	KycNotStarted                   string = "kyc_status_not_started"
	KycStatusInProgress             string = "kyc_status_in_progress"
	KycStatusInReview               string = "kyc_status_in_review"
	KycStatusAdditionalDataRequired string = "kyc_status_additional_data_required"
	KycStatusApproved               string = "kyc_status_approved"
	KycStatusDeclined               string = "kyc_status_declined"
	KycStatusBanned                 string = "kyc_status_banned"

	KycLevel0 int = 0
	KycLevel1 int = 1
	KycLevel2 int = 2
	KycLevel3 int = 3
	KycLevel4 int = 4
	KycLevel5 int = 5
)

type KycStatus struct {
	UserId                 int64          `json:"user_id,omitempty"`
	Status                 string         `json:"status,omitempty"`
	LevelApproved          int            `json:"level_approved,omitempty"`
	AdditionalDataRequired *[]KycDocument `json:"additional_data_required,omitempty"`
}

func NewKycStatus(userId int64) *KycStatus {
	return &KycStatus{
		UserId:                 userId,
		Status:                 KycNotStarted,
		LevelApproved:          KycLevel0,
		AdditionalDataRequired: nil,
	}
}
