package handler

import (
	"net/http"
	"strconv"
	"time"
)

func (ah *AdminHandler) LoginHandler(w http.ResponseWriter, r *http.Request) {

	username := r.FormValue("username")
	password := r.FormValue("password")

	admin, err := ah.serviceAdmin.ValidateAdmin(username, password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	cookie := &http.Cookie{
		Name:    "admin-session",
		Value:   strconv.Itoa(admin.ID),
		Path:    "/",
		Expires: time.Now().Add(1 * time.Hour),
	}

	http.SetCookie(w, cookie)
	http.Redirect(w, r, "/dashboard", http.StatusFound)
}
