package repository

import "day-24/model"

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
