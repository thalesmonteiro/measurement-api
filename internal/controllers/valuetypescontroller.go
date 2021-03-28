package controllers

import (
	"api/internal/models"
	"api/internal/utils"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

// @Tags Type Access
// @Summary Create a measurement type
// @Description Create a measurement type in the database
// @Success 200 "Success"
// @Accept json
// @Param userId body models.ValueTypes true "Inform userId and description"
// @Router /type/ [POST]
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

// @Tags Type Access
// @Summary Get Types by username
// @Description List all types by a username
// @Success 200 "Success"
// @Param username path string true "username"
// @Router /type{username}/ [GET]
func GetTypesByUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	username := vars["username"]
	types := models.GetTypesByUser(username)

	utils.RespondWithJson(w, http.StatusOK, types)
}

// @Tags Type Access
// @Summary Get Types in all users by description
// @Description List all types in all users by description
// @Success 200 "Success"
// @Param description path string true "description"
// @Router /alltypes/{description}/ [GET]
func GetTypeForAllUsersByDescription(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	description := vars["description"]
	types := models.GetTypeForAllUsersByDescription(description)
	utils.RespondWithJson(w, http.StatusOK, types)
}
