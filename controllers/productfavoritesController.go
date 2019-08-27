package controllers

import (
	"encoding/json"
	"go-contacts/models"
	u "go-contacts/utils"
	"net/http"
)

var AddList = func(w http.ResponseWriter, r *http.Request) {

	user := r.Context().Value("user").(uint) //Grab the id of the user that send the request
	model := &models.FavoritesProduct{}

	err := json.NewDecoder(r.Body).Decode(model)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}

	model.UserId = user
	resp := model.Create()
	u.Respond(w, resp)
}

var RemoveListFor = func(w http.ResponseWriter, r *http.Request) {

	model := &models.FavoritesProduct{}
	err := json.NewDecoder(r.Body).Decode(model) //decode the request body into struct and failed if any error occur
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}

	resp := models.Delete(model.ID)
	u.Respond(w, resp)

}
