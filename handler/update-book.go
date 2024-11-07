package handler

import (
	"database/sql"
	"day-24/library"
	"day-24/model"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

func (bh *BookHandler) UpdateBookHandler(w http.ResponseWriter, r *http.Request) {

	bookIDStr := chi.URLParam(r, "id")
	bookID, err := strconv.Atoi(bookIDStr)
	if err != nil || bookID <= 0 {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}

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

	coverPath := ""
	if _, _, err := r.FormFile("cover"); err == nil {
		coverPath, err = library.UploadFile(r, "cover", "./uploads/cover", "jpg")
		if err != nil {
			http.Error(w, fmt.Sprintf("Error uploading cover file: %v", err), http.StatusInternalServerError)
			return
		}
	}
	bookCover := sql.NullString{Valid: coverPath != "", String: coverPath}

	bookFilePath := ""
	if _, _, err := r.FormFile("file"); err == nil {
		bookFilePath, err = library.UploadFile(r, "file", "./uploads/books", "pdf")
		if err != nil {
			http.Error(w, fmt.Sprintf("Error uploading book file: %v", err), http.StatusInternalServerError)
			return
		}
	}
	bookFile := sql.NullString{Valid: bookFilePath != "", String: bookFilePath}

	book := model.Book{
		ID:        bookID,
		Title:     title,
		Category:  category,
		Author:    author,
		Price:     price,
		Discount:  discount,
		BookCover: bookCover,
		BookFile:  bookFile,
	}

	err = bh.serviceBooks.UpdateBook(book)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error updating book: %v", err), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/book-list", http.StatusSeeOther)
}
