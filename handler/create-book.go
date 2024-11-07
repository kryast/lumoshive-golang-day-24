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
		discount = 0
	}

	coverPath, err := library.UploadFile(r, "cover", "./uploads/cover", "jpg")
	if err != nil {
		http.Error(w, fmt.Sprintf("Error uploading cover file: %v", err), http.StatusInternalServerError)
		return
	}

	bookCover := sql.NullString{Valid: coverPath != "", String: coverPath}

	bookFilePath, err := library.UploadFile(r, "file", "./uploads/books", "pdf")
	if err != nil {
		http.Error(w, fmt.Sprintf("Error uploading book file: %v", err), http.StatusInternalServerError)
		return
	}

	bookFile := sql.NullString{Valid: bookFilePath != "", String: bookFilePath}

	book := model.Book{
		Title:     title,
		Category:  category,
		Author:    author,
		Price:     price,
		Discount:  discount,
		BookCover: bookCover,
		BookFile:  bookFile,
	}

	err = bh.serviceBooks.CreateBook(book)
	if err != nil {
		http.Error(w, "Error creating book: "+err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/book-list", http.StatusSeeOther)
}
