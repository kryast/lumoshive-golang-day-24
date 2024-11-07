package handler

import (
	"database/sql"
	"day-24/library"
	"day-24/model"
	"day-24/service"
	"fmt"
	"net/http"
	"strconv"
)

type BookHandler struct {
	serviceBooks service.BookService
}

func NewBookHandler(bs service.BookService) BookHandler {
	return BookHandler{serviceBooks: bs}
}

func (bh *BookHandler) CreateBookHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	// Ambil form data
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

	// Upload file cover, jika tidak ada, biarkan kosong (""), dan convert menjadi sql.NullString
	coverPath, err := library.UploadFile(r, "cover", "./uploads/cover", "jpg")
	if err != nil {
		http.Error(w, fmt.Sprintf("Error uploading cover file: %v", err), http.StatusInternalServerError)
		return
	}
	// Mengonversi coverPath menjadi sql.NullString
	bookCover := sql.NullString{Valid: coverPath != "", String: coverPath}

	// Proses file buku, jika tidak ada, biarkan kosong (""), dan convert menjadi sql.NullString
	bookFilePath, err := library.UploadFile(r, "file", "./uploads/books", "pdf")
	if err != nil {
		http.Error(w, fmt.Sprintf("Error uploading book file: %v", err), http.StatusInternalServerError)
		return
	}
	// Mengonversi bookFilePath menjadi sql.NullString
	bookFile := sql.NullString{Valid: bookFilePath != "", String: bookFilePath}

	// Simpan path file ke objek book
	book := model.Book{
		Title:     title,
		Category:  category,
		Author:    author,
		Price:     price,
		Discount:  discount,
		BookCover: bookCover, // Menyimpan path cover file (bisa kosong)
		BookFile:  bookFile,  // Menyimpan path file buku (bisa kosong)
	}

	// Simpan ke database melalui service dan repository
	err = bh.serviceBooks.CreateBook(book)
	if err != nil {
		http.Error(w, "Error creating book: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Tampilkan halaman sukses atau redirect ke halaman daftar buku
	http.Redirect(w, r, "/book-list", http.StatusSeeOther)
}
