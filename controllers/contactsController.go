package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"informationserver/models"
	u "informationserver/utils"
	"net/http"

	"strconv"
)

//создание контакта
var CreateContact = func(w http.ResponseWriter, r *http.Request) {

	user := r.Context().Value("user").(uint) //Получение идентификатора пользователя, отправившего запрос
	contact := &models.Contact{}

	err := json.NewDecoder(r.Body).Decode(contact)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}

	contact.UserId = user
	resp := contact.Create()
	u.Respond(w, resp)
}

var GetContactsFor = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		//Переданный параметр пути не является целым числом
		u.Respond(w, u.Message(false, "There was an error in your request"))
		return
	}
	data := models.GetContacts(uint(id))
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}

var GetContactsForCurentUser = func(w http.ResponseWriter, r *http.Request) {
	user_id, err := u.GetContextValueUint(r, "user")
	if err != nil {
		//Переданный параметр пути не является целым числом
		u.Respond(w, u.Message(false, "There was an error in your request"))
		return
	}
	data := models.GetContacts(uint(user_id))
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}
