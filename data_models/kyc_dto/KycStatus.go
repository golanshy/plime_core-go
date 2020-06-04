package kyc_dto

const (
	KycNotStarted                   string = "kyc_status_not_started"
	KycStatusInProgress             string = "kyc_status_in_progress"
	KycStatusInReview               string = "kyc_status_in_review"
	KycStatusAdditionalDataRequired string = "kyc_status_additional_data_required"
	KycStatusApproved               string = "kyc_status_kyc_status_approved"
	KycStatusDeclined               string = "kyc_status_declined"
	KycStatusBanned                 string = "kyc_status_banned"

	KycLevel0 int = 0
	KycLevel1 int = 1
	KycLevel2 int = 2
	KycLevel3 int = 3
	KycLevel4 int = 4
	KycLevel5 int = 5

	DocumentTypePassport       string = "document_type_passport"
	DocumentTypeDrivingLicence string = "document_type_driving_licence"
	DocumentTypePhotoId        string = "document_type_photo_id"
	DocumentTypeSelfie         string = "document_type_selfie"
	DocumentTypeSelfieVideo    string = "document_type_selfie_video"
	DocumentTypeProofOfAddress string = "document_type_proof_of_address"

	DocumentStatusRequired string = "document_status_required"
	DocumentStatusSupplied string = "document_status_supplied"
	DocumentStatusInReview string = "document_status_in_review"
	DocumentStatusApproved string = "document_status_approved"
	DocumentStatusRejected string = "document_status_rejected"
)

type KycStatus struct {
	Status        string        `json:"status,omitempty"`
	LevelApproved int           `json:"level_approved,omitempty"`
	Documents     []KycDocument `json:"documents,omitempty"`
}

type KycDocument struct {
	Name     string   `json:"name,omitempty"`
	Quantity int      `json:"quantity"`
	Types    []string `json:"types,omitempty"`
	Status   string   `json:"status,omitempty"`
	Required bool     `json:"required,omitempty"`
	Details  string   `json:"details,omitempty"`
}

func NewKycStatus() *KycStatus {
	return &KycStatus{
		Status:        KycNotStarted,
		LevelApproved: KycLevel0,
		Documents:     make([]KycDocument, 0),
	}
}

func NewKycDocument() *KycDocument {
	return &KycDocument{
		Quantity: 1,
		Status:   DocumentStatusRequired,
		Required: true,
		Types:    make([]string, 0),
	}
}
