package service

import "day-24/model"

func (bs *BookService) UpdateBook(book model.Book) error {
	return bs.RepoBook.UpdateBook(book)
}
