package date_utils

import (
	"github.com/golanshy/plime_core-go/utils/rest_errors"
	"time"
)

const (
	apiDateFormat  = "2006-01-02T15:04:05.000Z"
	apiSdLayout = "2006-01-02 15:04:05"
)

func GetNow() time.Time {
	return time.Now().UTC()
}

func GetNowString() string {
	return GetNow().Format(apiDateFormat)
}

func GetNowDBFormat() string {
	return GetNow().Format(apiSdLayout)
}

func FormatAPITime(value string) (*time.Time, *rest_errors.RestErr) {
	result, dateErr := time.Parse(apiDateFormat, value)
	if dateErr != nil {
		return nil, rest_errors.NewBadRequestError("invalid date format")
	}
	return &result, nil
}