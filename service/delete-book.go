package service

func (bs *BookService) DeleteBook(bookID int) error {
	// Panggil repository untuk menghapus buku
	err := bs.RepoBook.DeleteBook(bookID)
	return err
}
