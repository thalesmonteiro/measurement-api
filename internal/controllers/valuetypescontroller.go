package controllers

import (
	"api/internal/models"
	"api/internal/utils"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

var NewValueTypes models.ValueTypes

func CreateValueTypes(w http.ResponseWriter, r *http.Request) {
	createValueType := &models.ValueTypes{}
	utils.ParseBody(r, createValueType)

	//Verify if already exists type create for the user with the same description
	hasValue := models.GetTypeByIdAndDescription(createValueType.UserID, createValueType.Description)

	if hasValue.UserID == createValueType.UserID {
		msg := fmt.Sprintf("\"cannot create measurement type for this user because already exists a type with the specified description. typeID: %d", createValueType.TypeID)
		utils.RespondWithError(w, http.StatusBadRequest, msg)
	}

	typeValue := createValueType.CreateValueType()
	msg := fmt.Sprintf("Value with description: %s, create successfully", typeValue.Description)
	utils.RespondWithJson(w, http.StatusCreated, map[string]string{"message": msg})
}

func GetTypesByUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	username := vars["username"]
	types := models.GetTypesByUser(username)

	utils.RespondWithJson(w, http.StatusOK, types)
}

func GetTypeForAllUsersByDescription(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	description := vars["description"]
	types := models.GetTypeForAllUsersByDescription(description)
	utils.RespondWithJson(w, http.StatusOK, types)
}
