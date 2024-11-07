package handler

import (
	"day-24/model"
	"net/http"
)

func (bh *BookHandler) BookListHandler(w http.ResponseWriter, r *http.Request) {

	books, err := bh.serviceBooks.GetAllBooks()
	if err != nil {
		http.Error(w, "Error fetching books from database: "+err.Error(), http.StatusInternalServerError)
		return
	}

	data := struct {
		Books []model.Book
	}{
		Books: books,
	}

	err = templates.ExecuteTemplate(w, "book-list-view", data)
	if err != nil {
		http.Error(w, "Error rendering template: "+err.Error(), http.StatusInternalServerError)
	}
}
