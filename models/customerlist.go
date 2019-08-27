package models

import (
	u "go-contacts/utils"

	"github.com/jinzhu/gorm"
)

type CustomerList struct {
	gorm.Model
	Name       string `json:"name"`
	UserId     uint   `json:"user_id"`
	CustomerId uint   `json:"customer_id"`
}

func (customerlist *CustomerList) Validate() (map[string]interface{}, bool) {

	if customerlist.Name == "" {
		return u.Message(false, "Customer List name should be on the payload"), false
	}

	if customerlist.UserId <= 0 {
		return u.Message(false, "User is not recognized"), false
	}

	return u.Message(true, "success"), true
}

func (customerlist *CustomerList) Create() map[string]interface{} {

	if resp, ok := customerlist.Validate(); !ok {
		return resp
	}

	GetDB().Create(customerlist)

	resp := u.Message(true, "success")
	resp["customer"] = customerlist
	return resp
}
