package kyc_dto

const (
	KycNotStarted                   string = "kyc_status_not_started"
	KycStatusInProgress             string = "kyc_status_in_progress"
	KycStatusInReview               string = "kyc_status_in_review"
	KycStatusAdditionalDataRequired string = "kyc_status_additional_data_required"
	KycStatusApproved               string = "kyc_status_approved"
	KycStatusDeclined               string = "kyc_status_declined"
	KycStatusBanned                 string = "kyc_status_banned"

	KycLevel0 int64 = 0
	KycLevel1 int64 = 1
	KycLevel2 int64 = 2
	KycLevel3 int64 = 3
	KycLevel4 int64 = 4
	KycLevel5 int64 = 5
)

type KycStatus struct {
	UserId                 string         `json:"user_id"`
	Status                 string        `json:"status"`
	LevelApproved          int64         `json:"level_approved"`
	AdditionalDataRequired []KycDocument `json:"additional_data_required,omitempty"`
}

func NewKycStatus(userId string) KycStatus {
	return KycStatus{
		UserId:        userId,
		Status:        KycNotStarted,
		LevelApproved: KycLevel0,
	}
}
