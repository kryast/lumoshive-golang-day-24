package handler

import (
	"html/template"
	"net/http"
)

var templates = template.Must(template.ParseGlob("view/*.html"))

func FormRegist(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "register-view", nil)
}
func FormLogin(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "login-view", nil)
}
