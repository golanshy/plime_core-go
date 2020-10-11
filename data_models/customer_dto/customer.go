package customer_dto

import (
	"github.com/golanshy/plime_core-go/data_models/address_dto"
	"github.com/golanshy/plime_core-go/data_models/user_dto"
)

type Customer struct {
	Id            int64               `json:"id"`
	Name          string              `json:"name,omitempty"`
	Details       string              `json:"details,omitempty"`
	ContactPerson user_dto.Users      `json:"contact_person,omitempty"`
	Status        string              `json:"status,omitempty"`
	Address       address_dto.Address `json:"address,omitempty"`
	DateCreated   string              `json:"date_created,omitempty"`
}
