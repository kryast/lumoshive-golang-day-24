package repository

func (br *BookRepositoryDB) DeleteBook(bookID int) error {
	query := `DELETE FROM books WHERE id = $1`
	_, err := br.DB.Exec(query, bookID)
	return err
}
