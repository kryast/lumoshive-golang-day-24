package handler

import (
	"net/http"
	"time"
)

func (ah *AdminHandler) LogoutHandler(w http.ResponseWriter, r *http.Request) {
	cookie := &http.Cookie{
		Name:    "admin-session",
		Value:   "",
		Path:    "/",
		Expires: time.Now().Add(-1 * time.Hour),
	}

	http.SetCookie(w, cookie)

	http.Redirect(w, r, "/", http.StatusFound)
}
