package handler

import (
	"day-24/library"
	"day-24/model"
	"day-24/service"
	"encoding/json"
	"fmt"
	"net/http"
)

type AdminHandler struct {
	serviceAdmin service.AdminService
}

func NewAdminHandler(service service.AdminService) *AdminHandler {
	return &AdminHandler{serviceAdmin: service}
}

func (ah *AdminHandler) CreateAdminHandler(w http.ResponseWriter, r *http.Request) {
	var admin model.Admin

	err := json.NewDecoder(r.Body).Decode(&admin)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error decoding request body: %v", err), http.StatusBadRequest)
		return
	}

	err = ah.serviceAdmin.CreateAdmin(admin)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error creating admin: %v", err), http.StatusInternalServerError)
		return
	}

	library.SuccessResponse(w, "Admin created successfully", admin)
}
