package handler

import (
	"html/template"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

var templates = template.Must(template.ParseGlob("view/*.html"))

func FormCreateBook(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "create-book-view", nil)
}
func (bh *BookHandler) FormEditBook(w http.ResponseWriter, r *http.Request) {
	bookIDStr := chi.URLParam(r, "id")
	bookID, err := strconv.Atoi(bookIDStr)
	if err != nil || bookID <= 0 {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}

	book, err := bh.serviceBooks.GetBookByID(bookID)
	if err != nil {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}

	templates.ExecuteTemplate(w, "edit-book-view", book)
}

func FormLogin(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "login-view", nil)
}

func Logout(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "logout-view", nil)
}
