package user_device_dto

import (
	"github.com/golanshy/plime_core-go/utils/rest_errors"
	"strings"
)

type UserDevice struct {
	Platform  string `json:"platform"`
	Make      string `json:"make"`
	Model     string `json:"model"`
	OsVersion string `json:"os_version"`
	AppName   string `json:"app_name"`
	PushId    string `json:"push_id"`
	UserId    string  `json:"user_id"`
}

func (d *UserDevice) Validate() *rest_errors.RestErr {
	if strings.TrimSpace(d.PushId) == "" {
		return rest_errors.NewBadRequestError("missing device push id")
	}
	return nil
}

type UserDevices struct {
	Results []UserDevice `json:"user_devices"`
}


