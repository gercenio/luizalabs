package models

import (
	"fmt"
	u "go-contacts/utils"

	"github.com/jinzhu/gorm"
)

type Customer struct {
	gorm.Model
	Name             string             `json:"name"`
	Email            string             `json:"email"`
	UserId           uint               `json:"user_id"`
	CustomerList     []CustomerList     `gorm:"foreignkey:CustomerId"`
	FavoritesProduct []FavoritesProduct `gorm:"foreignkey:CustomerId;foreignkey:CustomerListId"`
}

/*
returns message and true if the requirement is met
*/
func (customer *Customer) Validate() (map[string]interface{}, bool) {

	if customer.Name == "" {
		return u.Message(false, "Contact name should be on the payload"), false
	}

	if customer.Email == "" {
		return u.Message(false, "Email should be on the payload"), false
	}

	if customer.UserId <= 0 {
		return u.Message(false, "User is not recognized"), false
	}

	//All the required parameters are present
	return u.Message(true, "success"), true
}

func (customer *Customer) Create() map[string]interface{} {

	if resp, ok := customer.Validate(); !ok {
		return resp
	}

	err := GetCustomerByEmail(customer.Email)
	if err != nil {

		resp := u.Message(false, "Email already registered")
		resp["customer"] = err
		return resp
	}

	GetDB().Create(customer)

	resp := u.Message(true, "success")
	resp["customer"] = customer
	return resp
}

func GetCustomer(id uint) *Customer {

	customer := &Customer{}
	err := GetDB().Table("customers").Where("id = ?", id).First(customer).Error
	if err != nil {
		return nil
	}
	return customer
}

func GetCustomerByEmail(email string) *Customer {

	customer := &Customer{}
	err := GetDB().Table("customers").Where("email = ?", email).First(customer).Error
	if err != nil {
		return nil
	}
	return customer
}

func DeleteCustomer(id uint) map[string]interface{} {

	err := GetDB().Table("customers").Where("id =?", id).Delete(&Customer{}).Error
	if err != nil {
		return u.Message(false, "Error")
	}
	return u.Message(true, "success")
}

func GetCustomers(user uint) []*Customer {

	customers := make([]*Customer, 0)
	err := GetDB().Table("customers").Where("user_id = ?", user).Find(&customers).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return customers
}
