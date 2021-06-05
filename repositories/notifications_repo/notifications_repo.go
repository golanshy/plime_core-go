package notifications_repo

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/golanshy/plime_core-go/data_models/notifications_dto"
	"github.com/golanshy/plime_core-go/data_models/user_dto"
	"github.com/golanshy/plime_core-go/logger"
	"github.com/golanshy/plime_core-go/utils/rest_errors"
	"os"
	"strings"
)

var (
	client *resty.Client
	apiUrl string
)

func init() {
	client = resty.New()
	apiUrl = strings.TrimSpace(os.Getenv("NOTIFICATIONS_API_URL"))
	if apiUrl == "" {
		panic(errors.New("missing NOTIFICATIONS_API_URL"))
	}
}

type RestNotificationsRepository interface {
	NotifyUser(authorization string, messageRequest notifications_dto.NotificationRequest) *rest_errors.RestErr
}

type notificationsRepository struct {
}

func NewRestNotificationsRepository() RestNotificationsRepository {
	return &notificationsRepository{}
}

func (r *notificationsRepository) NotifyUser(authorization string, messageRequest notifications_dto.NotificationRequest) *rest_errors.RestErr {
	headers := make(map[string]string)
	headers["Authorization"] = authorization
	headers["Content-Type"] = "application/json"
	response, err := client.R().
		SetHeaders(headers).
		SetBody(messageRequest).
		Post(fmt.Sprintf("%s/notifications", apiUrl))

	if err != nil {
		logger.Error("error sending notification to user", err)
		return  rest_errors.NewInternalServerError("error sending notification to user", err)
	}
	if response.StatusCode() > 299 {
		restErr, err := rest_errors.NewRestErrorFromBytes(response.Body())
		if err != nil {
			logger.Error("invalid error interface when trying to send notification to user", err)
			return rest_errors.NewInternalServerError("invalid error interface when trying to send notification to user", err)
		}
		restErr.Status = response.StatusCode()
		return  restErr
	}
	var user user_dto.User
	if err := json.Unmarshal(response.Body(), &user); err != nil {
		logger.Error("error unmarshaling json response when trying to send notification to user", err)
		return rest_errors.NewInternalServerError("error unmarshaling json response when trying to send notification to user", err)
	}
	return nil
}
