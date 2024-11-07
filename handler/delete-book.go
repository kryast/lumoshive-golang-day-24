package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

func (bh *BookHandler) DeleteBookHandler(w http.ResponseWriter, r *http.Request) {
	bookIDStr := chi.URLParam(r, "id")
	bookID, err := strconv.Atoi(bookIDStr)
	if err != nil || bookID <= 0 {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}

	err = bh.serviceBooks.DeleteBook(bookID)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error deleting book: %v", err), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/book-list", http.StatusSeeOther)
}
