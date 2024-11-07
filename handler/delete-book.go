package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

func (bh *BookHandler) DeleteBookHandler(w http.ResponseWriter, r *http.Request) {
	// Ambil ID buku yang akan dihapus dari URL
	bookIDStr := chi.URLParam(r, "id") // Gunakan chi.URLParam untuk mengambil parameter dari URL
	bookID, err := strconv.Atoi(bookIDStr)
	if err != nil || bookID <= 0 {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}

	// Panggil service untuk menghapus buku berdasarkan ID
	err = bh.serviceBooks.DeleteBook(bookID)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error deleting book: %v", err), http.StatusInternalServerError)
		return
	}

	// Redirect ke halaman daftar buku setelah sukses
	http.Redirect(w, r, "/book-list", http.StatusSeeOther)
}
