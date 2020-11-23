package notifications_dto

import (
	"github.com/golanshy/plime_core-go/utils/rest_errors"
	"strings"
)

type NotificationRequest struct {
	Channel      string   `json:"channel"`
	UserIds      []string `json:"user_ids"`
	Destinations []string `json:"destinations,omitempty"`
	Subject      string   `json:"subject,omitempty"`
	HtmlMessage  string   `json:"html_message,omitempty"`
	Message      string   `json:"message,omitempty"`
	Payload      string   `json:"payload,omitempty"`
	DateCreated  string   `json:"date_created"`
}

const (
	ChannelEmail    string = "email"
	ChannelText     string = "text"
	ChannelWhatsApp string = "whatsapp"
	ChannelPush     string = "push"
)

func (r *NotificationRequest) Validate() *rest_errors.RestErr {
	if len(r.UserIds) == 0 && len(r.Destinations) == 0 {
		return rest_errors.NewBadRequestError("invalid user ids or destinations")
	}
	if strings.TrimSpace(r.Message) == "" && strings.TrimSpace(r.Channel) == "" {
		return rest_errors.NewBadRequestError("invalid channel")
	}
	if strings.TrimSpace(r.Message) == "" && strings.TrimSpace(r.Payload) == "" {
		return rest_errors.NewBadRequestError("invalid message or payload")
	}
	return nil
}
