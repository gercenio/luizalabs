package models

import (
	u "go-contacts/utils"

	"github.com/jinzhu/gorm"
)

type FavoritesProduct struct {
	gorm.Model
	ProductId      uint `json:"product_id"`
	UserId         uint `json:"user_id"`
	CustomerId     uint `json:"customer_id"`
	CustomerListId uint `json:"customerlist_id"`
}

func (favoritesproduct *FavoritesProduct) Validate() (map[string]interface{}, bool) {

	if favoritesproduct.ProductId <= 0 {
		return u.Message(false, "Product is not invalid"), false
	}

	if favoritesproduct.CustomerId <= 0 {
		return u.Message(false, "Customer is not invalid"), false
	}

	if favoritesproduct.CustomerListId <= 0 {
		return u.Message(false, "Customer List is not invalid"), false
	}

	if favoritesproduct.UserId <= 0 {
		return u.Message(false, "User is not recognized"), false
	}

	return u.Message(true, "success"), true
}

func (favoriteproduct *FavoritesProduct) Create() map[string]interface{} {

	if resp, ok := favoriteproduct.Validate(); !ok {
		return resp
	}

	GetDB().Create(favoriteproduct)

	resp := u.Message(true, "success")
	resp["customer"] = favoriteproduct
	return resp
}

func Delete(id uint) map[string]interface{} {

	err := GetDB().Table("favorites_products").Where("id =?", id).Delete(&FavoritesProduct{}).Error
	if err != nil {
		return u.Message(false, "Error")
	}
	return u.Message(true, "success")
}
