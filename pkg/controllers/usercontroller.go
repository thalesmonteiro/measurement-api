package controllers

import (
	"api/pkg/models"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	//"github.com/thalesmonteiro/measurementApi/pkg/models"
	//"github.com/thalesmonteiro/measurementApi/pkg/utils"
	"api/pkg/utils"
	"net/http"
	"strconv"
)

var NewUser models.User

// Username in request body
func CreateUser(w http.ResponseWriter, r *http.Request) {
	//Parse the body to user model
	createUser := &models.User{}
	utils.ParseBody(r, createUser)

	//Verify if the userID already exists
	hasUser := models.GetUserByUsername(createUser.Username)
	if hasUser.Username == createUser.Username {
		utils.RespondWithError(w, http.StatusBadRequest, "Cannot create a user, username already exist")
		return
	}

	//Create User
	user := createUser.CreateUser()

	//Found the userId created in database
	userF := models.GetUserByUsername(createUser.Username)
	user.UserID = userF.UserID

	utils.RespondWithJson(w, http.StatusCreated, user)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	users := models.GetAllUser()

	utils.RespondWithJson(w, http.StatusOK, users)
}

func GetUserById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["userId"]
	ID, err := strconv.ParseInt(userId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	userDetails, _ := models.GetUserById(ID)
	res, _ := json.Marshal(userDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	var updateUser = &models.User{}
	utils.ParseBody(r, updateUser)
	vars := mux.Vars(r)
	userId := vars["userId"]
	ID, err := strconv.ParseInt(userId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	userDetails, db := models.GetUserById(ID)
	if updateUser.Username != "" {
		userDetails.Username = updateUser.Username
	}

	db.Save(&userDetails)
	res, _ := json.Marshal(userDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["userId"]
	ID, err := strconv.ParseInt(userId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	user := models.DeleteUser(ID)
	res, _ := json.Marshal(user)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
