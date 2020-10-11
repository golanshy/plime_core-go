package role_dto

const (
	PlimeAdmin     string = "plime_admin"
	PlimeUser      string = "plime_user"
	PlimeModerator string = "plime_moderator"
	PlimeViewer    string = "plime_viewer"
	Admin          string = "admin"
	Moderator      string = "moderator"
	Viewer         string = "viewer"
)

type Role struct {
	Name string `json:"name,omitempty"`
}
