package repository

func (br *BookRepositoryDB) DeleteBook(bookID int) error {
	// Menjalankan query SQL untuk menghapus buku berdasarkan ID
	query := `DELETE FROM books WHERE id = $1`
	_, err := br.DB.Exec(query, bookID)
	return err
}
