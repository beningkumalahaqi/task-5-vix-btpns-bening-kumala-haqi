package helpers

import (
	"encoding/json"
	"net/http"
)

type ResoponseWithData struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

type ResoponseWithoutData struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func Response(w http.ResponseWriter, code int, message string, payload any) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	var response any
	status := "success"

	if code >= 400 {
		status = "failed"
	}

	if payload != nil {
		response = &ResoponseWithData{
			Status: status,
			Message: message,
			Data: payload,
		}
	} else {
		response = &ResoponseWithoutData{
			Status: status,
			Message: message,
		}
	}


	res, _ := json.Marshal(response)
	w.Write(res)
}