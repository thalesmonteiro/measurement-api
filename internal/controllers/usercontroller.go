package controllers

import (
	"api/internal/models"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	//"github.com/thalesmonteiro/measurementApi/internal/models"
	//"github.com/thalesmonteiro/measurementApi/internal/utils"
	"api/internal/utils"
	"net/http"
	"strconv"
)

// @Tags User Access
// @Summary Create a User
// @Description Create a user in the database
// @Success 200 "Success"
// @Success 204 "No Content"
// @Accept json
// @Param username body models.User true "username"
// @Router /user/ [POST]
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

// @Tags User Access
// @Summary Get all Users
// @Description List all users in database
// @Success 200 "Success"
// @Router /user/ [GET]
func GetAllUser(w http.ResponseWriter, r *http.Request) {
	users := models.GetAllUser()

	utils.RespondWithJson(w, http.StatusOK, users)
}

// TODO validar esse endpoint
func GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	parameter := vars["userInfo"]
	ID, err := strconv.ParseInt(parameter, 0, 0)
	if err != nil {
		userDetails := models.GetUserByUsername(parameter)
		res, _ := json.Marshal(userDetails)

		if userDetails.UserID == 0 {
			utils.RespondWithError(w, http.StatusBadRequest, "There is no registered user with the given username.")
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(res)

		return
	}
	userDetails, _ := models.GetUserById(ID)
	if userDetails.UserID == 0 {
		utils.RespondWithError(w, http.StatusBadRequest, "There is no registered user with the given userId.")
		return
	}
	res, _ := json.Marshal(userDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// @Tags User Access
// @Summary Get an User
// @Description List of a user in database
// @Success 200 "Success"
// @Success 204 "No Content"
// @Accept json
// @Param username body models.User true "username"
// @Router /user/ [GET]
func GetUserByUsername(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	username := vars["userName"]
	//ID, err := strconv.ParseInt(userId, 0, 0)
	//if err != nil {
	//	fmt.Println("Error while parsing")
	//}
	userDetails := models.GetUserByUsername(username)
	res, _ := json.Marshal(userDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// @Tags User Access
// @Summary Update User by ID
// @Description Update an user information
// @Success 200 "Success"
// @Success 204 "No Content"
// @Accept json
// @Param username and id body models.User true "username"
// @Param userId path int true "userId"
// @Router /user{userId}/ [PUT]
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

// @Tags User Access
// @Summary Delete User by ID
// @Description Update an user information
// @Success 200 "Success"
// @Accept json
// @Param userId path int true "userId"
// @Router /user{userId}/ [DELETE]
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

// @Tags User Access
// @Summary Get Users has measures
// @Description Update an user information
// @Success 200 "Success"
// @Accept json
// @Router /usermeasure/ [GET]
func GetUsersHasMeasure(w http.ResponseWriter, r *http.Request) {
	users := models.GetUsersHasMeasure()
	utils.RespondWithJson(w, http.StatusOK, users)
}
