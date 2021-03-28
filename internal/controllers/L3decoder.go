package controllers

import (
	"api/internal/models"
	"api/internal/utils"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// @Tags Measurement Request
// @Summary L3 Communication
// @Description Communicate with ESP 32 using L3 protocol.
// @Success 200 "Success"
// @Success 204 "No Content"
// @Accept json
// @Param Data path string true "String with L3 Packet"
// @Router /measure/ [POST]
func L3DecoderPayload(w http.ResponseWriter, r *http.Request) {
	//paylod with 12 bytes
	payload := &models.L3{}
	utils.ParseBody(r, payload)
	data := utils.AsByteSlice(payload.Data)

	//byte 3 is the code instruction
	messageCode := data[3]
	//byte 5 to 9 is the dataPadding
	//dataPadding := data[4:10]
	//2 bytes of data
	//value := dataPadding[1:3]

	//Esp Address
	url := fmt.Sprintf("localhost:8080?code=%s", messageCode)

	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", body)

	CreateMeasure(w, r, string(body))

	fmt.Print(data)
	return
}
