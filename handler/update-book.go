package handler

import (
	"day-24/library"
	"day-24/model"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

// handler/book_handler.go
// handler/book_handler.go

func (bh *BookHandler) UpdateBookHandler(w http.ResponseWriter, r *http.Request) {
	// Ambil ID dari URL
	bookIDStr := chi.URLParam(r, "id") // Ambil ID dari URL
	bookID, err := strconv.Atoi(bookIDStr)
	if err != nil || bookID <= 0 {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}

	// Ambil data form lainnya
	title := r.FormValue("title")
	category := r.FormValue("category")
	author := r.FormValue("author")
	priceStr := r.FormValue("price")
	price, err := strconv.ParseFloat(priceStr, 64)
	if err != nil {
		http.Error(w, "Invalid price format", http.StatusBadRequest)
		return
	}

	discountStr := r.FormValue("discount")
	discount, err := strconv.ParseFloat(discountStr, 64)
	if err != nil {
		discount = 0 // Jika diskon tidak ada, set ke 0
	}

	// Cek jika ada file cover yang diupload
	coverPath := ""
	if _, _, err := r.FormFile("cover"); err == nil { // Hanya jika file cover ada
		coverPath, err = library.UploadFile(r, "cover", "./uploads/cover", "jpg")
		if err != nil {
			http.Error(w, fmt.Sprintf("Error uploading cover file: %v", err), http.StatusInternalServerError)
			return
		}
	}

	// Cek jika ada file buku yang diupload
	bookFilePath := ""
	if _, _, err := r.FormFile("file"); err == nil { // Hanya jika file buku ada
		bookFilePath, err = library.UploadFile(r, "file", "./uploads/books", "pdf")
		if err != nil {
			http.Error(w, fmt.Sprintf("Error uploading book file: %v", err), http.StatusInternalServerError)
			return
		}
	}

	// Membuat objek book dengan data yang ada
	book := model.Book{
		ID:        bookID,
		Title:     title,
		Category:  category,
		Author:    author,
		Price:     price,
		Discount:  discount,
		BookCover: coverPath,    // Menyimpan path cover file jika ada
		BookFile:  bookFilePath, // Menyimpan path file buku jika ada
	}

	// Panggil service untuk memperbarui data buku
	err = bh.serviceBooks.UpdateBook(book)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error updating book: %v", err), http.StatusInternalServerError)
		return
	}

	// Redirect ke daftar buku
	http.Redirect(w, r, "/book-list", http.StatusSeeOther)
}
