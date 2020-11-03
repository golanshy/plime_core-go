package role_dto

const (
	PlimeAdmin        string = "plime_admin"
	PlimeUser         string = "plime_user"
	PlimeModerator    string = "plime_moderator"
	PlimeViewer       string = "plime_viewer"
	CustomerAdmin     string = "customer_admin"
	CustomerModerator string = "customer_moderator"
	CustomerViewer    string = "customer_viewer"
)

type UserRole struct {
	UserId                int64  `json:"user_id"`
	CustomerId            string `json:"customer_id"`
	CustomerName          string `json:"customer_name"`
	UserEmail             string `json:"user_email,omitempty"`
	UserFirstName         string `json:"user_first_name,omitempty"`
	UserLastName          string `json:"user_last_name,omitempty"`
	CompanyName           string `json:"company_name,omitempty"`
	CompanyRegisteredName string `json:"company_registered_name,omitempty"`
	CompanyRegisteredId   string `json:"company_registered_id,omitempty"`
	Role                  string `json:"role,omitempty"`
}
