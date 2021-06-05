package notifications_dto

import (
	"github.com/golanshy/plime_core-go/utils/rest_errors"
	"strings"
)

type MessageContent struct {
	MessageTitle     string `json:"message_title"`
	MessageSubTitle  string `json:"message_sub_title"`
	MessagePart1     string `json:"message_part_1"`
	MessagePart2     string `json:"message_part_2"`
	MessagePart3     string `json:"message_part_3"`
	BodyImage        string `json:"body_image"`
	BodyImageUrl     string `json:"body_image_url"`
	ActionButtonText string `json:"action_button_text"`
	ActionButtonURL  string `json:"action_button_url"`
}

type NotificationRequest struct {
	Channel            string         `json:"channel"`
	UserIds            []string       `json:"user_ids"`
	Destinations       []string       `json:"destinations,omitempty"`
	Subject            string         `json:"subject,omitempty"`
	HtmlMessage        string         `json:"html_message,omitempty"`
	HtmlMessageContent MessageContent `json:"html_message_content"`
	Message            string         `json:"message,omitempty"`
	Payload            string         `json:"payload,omitempty"`
	DateCreated        string         `json:"date_created"`
}

const (
	ChannelEmail    string = "email"
	ChannelText     string = "text"
	ChannelWhatsApp string = "whatsapp"
	ChannelPush     string = "push"
)

func (r *NotificationRequest) Validate() *rest_errors.RestErr {
	if len(r.UserIds) == 0 && len(r.Destinations) == 0 {
		return rest_errors.NewBadRequestError("missing user ids or destinations")
	}

	if strings.TrimSpace(r.Channel) != "" {
		return rest_errors.NewBadRequestError("missing channel")
	}

	if strings.TrimSpace(r.Channel) == ChannelEmail &&
		strings.TrimSpace(r.Channel) == ChannelText &&
		strings.TrimSpace(r.Channel) == ChannelWhatsApp &&
		strings.TrimSpace(r.Channel) == ChannelPush {
		return rest_errors.NewBadRequestError("invalid channel")
	}

	return nil
}
