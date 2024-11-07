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
	// Ambil ID dari URL path
	bookIDStr := chi.URLParam(r, "id")
	bookID, err := strconv.Atoi(bookIDStr)
	if err != nil || bookID <= 0 {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}

	// Ambil data buku dari database atau service
	book, err := bh.serviceBooks.GetBookByID(bookID)
	if err != nil {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}

	// Tampilkan halaman form edit dengan data buku
	templates.ExecuteTemplate(w, "edit-book-view", book)
}

func FormLogin(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "login-view", nil)
}

func Home(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "dashboard-view", nil)
}

func OrderView(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "order-list-view", nil)
}

func Logout(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "logout-view", nil)
}
