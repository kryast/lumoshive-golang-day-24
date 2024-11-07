package handler

import (
	"html/template"
	"net/http"
)

var templates = template.Must(template.ParseGlob("view/*.html"))

func FormCreateBook(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "create-book-view", nil)
}
func FormLogin(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "login-view", nil)
}

func Home(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "dashboard-view", nil)
}
