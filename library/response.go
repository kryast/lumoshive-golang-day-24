package library

import (
	"day-24/model"
	"encoding/json"
	"net/http"
)

func BadResponse(w http.ResponseWriter, message string) {
	badResponse := model.Response{
		StatusCode: http.StatusBadRequest,
		Message:    "Error server",
		Data:       nil,
	}
	json.NewEncoder(w).Encode(badResponse)
}

func SuccessResponse(w http.ResponseWriter, message string, data any) {
	badResponse := model.Response{
		StatusCode: http.StatusBadRequest,
		Message:    message,
		Data:       data,
	}
	json.NewEncoder(w).Encode(badResponse)
}
