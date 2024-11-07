package repository

import (
	"database/sql"
	"day-24/model"
)

type BookRepositoryDB struct {
	DB *sql.DB
}

func NewBookRepository(db *sql.DB) BookRepositoryDB {
	return BookRepositoryDB{DB: db}
}

func (r *BookRepositoryDB) CreateDataBook(book model.Book) error {
	query := "INSERT INTO books (title , category, author, price, discount, book_cover, book_file) VALUES ($1 , $2 , $3, $4, $5, $6 , $7) RETURNING id"

	err := r.DB.QueryRow(query, book.Title, book.Category, book.Author, book.Price, book.Discount, book.BookCover, book.BookFile).Scan(&book.ID)
	if err != nil {
		return err
	}

	return nil
}
