package controllers

import (
	"api/internal/models"
	"api/internal/utils"
	"github.com/gorilla/mux"
	"net/http"
)

//https://stackoverflow.com/questions/48646580/convert-a-bitstring-into-a-byte-array

//Expected
//{
//"typeId": 1,
//"value": "010110111001"
//}

func CreateMeasure(w http.ResponseWriter, r *http.Request, value string) {
	createMeasure := &models.Measure{}
	utils.ParseBody(r, createMeasure)
	createMeasure.Value = value

	hasType := models.GetTypeByID(createMeasure.TypeID)

	if hasType.TypeID == 0 {
		utils.RespondWithError(w, http.StatusBadRequest, "Cannot create measure. TypeID does not exist")
		return
	}
	_ = createMeasure.CreateMeasure()

	msg := "Measure create successfully."
	utils.RespondWithJson(w, http.StatusCreated, map[string]string{"message": msg})
}

func GetAllMeasureFromUsername(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	username := vars["username"]
	measure := models.GetAllMeasureFromUsername(username)

	utils.RespondWithJson(w, http.StatusOK, measure)
}
