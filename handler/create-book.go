package handler

import (
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
	bookName := r.FormValue("bookName")
	bookType := r.FormValue("bookType")
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

	coverPath, err := library.UploadFile(r, "cover", "./uploads/cover", "jpg")
	if err != nil {
		http.Error(w, fmt.Sprintf("Error uploading cover file: %v", err), http.StatusInternalServerError)
		return
	}

	// Proses file buku dengan fungsi uploadFile
	bookFilePath, err := library.UploadFile(r, "file", "./uploads/books", "pdf")
	if err != nil {
		http.Error(w, fmt.Sprintf("Error uploading book file: %v", err), http.StatusInternalServerError)
		return
	}

	// Simpan path file ke objek book
	book := model.Book{
		Title:     bookName,
		Category:  bookType,
		Author:    author,
		Price:     price,
		Discount:  discount,
		BookCover: coverPath,    // Menyimpan path cover file
		BookFile:  bookFilePath, // Menyimpan path file buku
	}

	// Simpan ke database melalui service dan repository
	err = bh.serviceBooks.CreateBook(book)
	if err != nil {
		http.Error(w, "Error creating book: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Tampilkan halaman sukses
	templates.ExecuteTemplate(w, "dashboard-view", book)
}
