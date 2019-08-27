package controllers

import (
	"encoding/json"
	"go-contacts/models"
	u "go-contacts/utils"
	"net/http"
)

var CreateCustomer = func(w http.ResponseWriter, r *http.Request) {

	user := r.Context().Value("user").(uint) //Grab the id of the user that send the request
	customer := &models.Customer{}

	err := json.NewDecoder(r.Body).Decode(customer)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}

	customer.UserId = user
	resp := customer.Create()
	u.Respond(w, resp)
}

var GetCustomersFor = func(w http.ResponseWriter, r *http.Request) {

	id := r.Context().Value("user").(uint)
	data := models.GetCustomers(id)
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}

var DeteleCustomerFor = func(w http.ResponseWriter, r *http.Request) {

	customer := &models.Customer{}
	err := json.NewDecoder(r.Body).Decode(customer) //decode the request body into struct and failed if any error occur
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}

	resp := models.DeleteCustomer(customer.ID)
	u.Respond(w, resp)

}
