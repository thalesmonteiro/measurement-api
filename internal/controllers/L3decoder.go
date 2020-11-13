package controllers

import (
	"api/internal/models"
	"api/internal/utils"
	"fmt"
	"net/http"
)

func L3DecoderPayload(w http.ResponseWriter, r *http.Request) {
	payload := &models.L3{}
	utils.ParseBody(r, payload)
	data := utils.AsByteSlice(payload.Data)

	//messageCode := data[3]
	//dataPadding := data[4:10]
	//value := dataPadding[1:3]
	//
	////Request a measure to sensor
	//if messageCode == 84 {
	//
	//}else{
	//	//Response sensor
	//	if messageCode == 86 {
	//		CreateMeasure(w, r, string(value))
	//	}else{
	//		return utils.RespondWithError()
	//	}
	//}

	fmt.Print(data)
	return
}
