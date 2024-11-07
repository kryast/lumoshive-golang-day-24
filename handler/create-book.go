package handler

import (
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

	// Ambil nilai form dari request
	title := r.FormValue("title")
	author := r.FormValue("author")
	category := r.FormValue("category")

	priceStr := r.FormValue("price")
	price, _ := strconv.ParseFloat(priceStr, 64)
	discountStr := r.FormValue("discount")
	discount, _ := strconv.ParseFloat(discountStr, 64)

	bookCover := r.FormValue("book_cover") // Contoh jika Anda ingin menangani upload cover
	bookFile := r.FormValue("book_file")   // Contoh jika Anda ingin menangani upload file

	// Buat objek Book dengan data yang diambil dari form
	// Pastikan data seperti price dan discount di-convert ke tipe yang sesuai jika perlu
	book := model.Book{
		Title:     title,
		Author:    author,
		Category:  category,
		Price:     price,
		Discount:  discount,
		BookCover: bookCover,
		BookFile:  bookFile,
	}

	// Panggil service untuk menyimpan data buku
	// Jika ada kesalahan saat menyimpan, tampilkan pesan error
	err := bh.serviceBooks.CreateBook(book)
	if err != nil {
		fmt.Println("error ", err)
		return
	}

}
