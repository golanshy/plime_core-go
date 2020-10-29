package kyb_dto

const (
	DataTypeString             string = "data_type_string"
	DataTypeInt             string = "data_type_int"
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

type KybDocument struct {
	Name     string   `json:"name,omitempty"`
	Types    []string `json:"types,omitempty"`
	Required bool     `json:"required,omitempty"`
	Details  string   `json:"details,omitempty"`
}

type KybDocuments struct {
	Documents []KybDocument `json:"documents,omitempty"`
}

func NewKybDocuments() *KybDocuments {
	newKyb := &KybDocuments{
		Documents: make([]KybDocument, 0),
	}
	typeId := make([]string, 0)
	typeSelfie := make([]string, 0)
	typeVideoSelfie := make([]string, 0)
	typeProofOfAddress := make([]string, 0)

	typeId = append(typeId, DocumentTypePassport)
	typeId = append(typeId, DocumentTypeDrivingLicence)
	typeId = append(typeId, DocumentTypePhotoId)

	typeSelfie = append(typeSelfie, DocumentTypeSelfie)
	typeVideoSelfie = append(typeVideoSelfie, DocumentTypeSelfieVideo)
	typeProofOfAddress = append(typeProofOfAddress, DocumentTypeProofOfAddress)

	newKyb.Documents = append(newKyb.Documents, KybDocument{
		Name:     "Passport, Driving licence or Id card",
		Types:    typeId,
		Required: true,
		Details:  "",
	})
	newKyb.Documents = append(newKyb.Documents, KybDocument{
		Name:     "Selfie",
		Types:    typeSelfie,
		Required: true,
		Details:  "",
	})
	newKyb.Documents = append(newKyb.Documents, KybDocument{
		Name:     "Video selfie",
		Types:    typeVideoSelfie,
		Required: true,
		Details:  "",
	})
	newKyb.Documents = append(newKyb.Documents, KybDocument{
		Name:     "Proof of address",
		Types:    typeProofOfAddress,
		Required: false,
		Details:  "",
	})

	return newKyb
}
