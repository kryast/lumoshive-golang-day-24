package repository

import (
	"database/sql"
	"day-24/model"
	"fmt"
)

func (r *BookRepositoryDB) GetAllBooks() ([]model.Book, error) {
	var books []model.Book
	rows, err := r.DB.Query("SELECT id, title, category, author, price, discount FROM books")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var book model.Book
		err := rows.Scan(&book.ID, &book.Title, &book.Category, &book.Author, &book.Price, &book.Discount)
		if err != nil {
			return nil, err
		}
		books = append(books, book)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return books, nil
}

func (br *BookRepositoryDB) FindByID(id int) (*model.Book, error) {
	query := `SELECT id, title, category, author, price, discount, book_cover, book_file 
	          FROM books WHERE id = $1`

	row := br.DB.QueryRow(query, id)

	book := &model.Book{}
	err := row.Scan(&book.ID, &book.Title, &book.Category, &book.Author, &book.Price, &book.Discount, &book.BookCover, &book.BookFile)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("book not found")
		}
		return nil, fmt.Errorf("could not scan book: %v", err)
	}
	return book, nil
}
