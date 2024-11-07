package handler

import (
	"day-24/model"
	"net/http"
)

// Handler untuk menampilkan daftar buku
func (bh *BookHandler) BookListHandler(w http.ResponseWriter, r *http.Request) {
	// Ambil data buku dari repository
	books, err := bh.serviceBooks.GetAllBooks()
	if err != nil {
		http.Error(w, "Error fetching books from database: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Kirim data buku ke template
	data := struct {
		Books []model.Book
	}{
		Books: books,
	}

	// Render template dengan data buku
	err = templates.ExecuteTemplate(w, "book-list-view", data)
	if err != nil {
		http.Error(w, "Error rendering template: "+err.Error(), http.StatusInternalServerError)
	}
}
