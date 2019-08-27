package controllers

import (
	"encoding/json"
	"go-contacts/models"
	u "go-contacts/utils"
	"net/http"
)

var CreateList = func(w http.ResponseWriter, r *http.Request) {

	user := r.Context().Value("user").(uint) //Grab the id of the user that send the request
	model := &models.CustomerList{}

	err := json.NewDecoder(r.Body).Decode(model)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}

	model.UserId = user
	resp := model.Create()
	u.Respond(w, resp)
}
