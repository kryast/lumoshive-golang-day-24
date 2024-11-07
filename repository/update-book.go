package repository

import "day-24/model"

func (r *BookRepositoryDB) UpdateBook(book model.Book) error {
	query := `
    UPDATE books
    SET title = $1, category = $2, author = $3, price = $4, discount = $5, book_cover = $6, book_file = $7
    WHERE id = $8`

	_, err := r.DB.Exec(query, book.Title, book.Category, book.Author, book.Price, book.Discount, book.BookCover, book.BookFile, book.ID)
	return err
}
