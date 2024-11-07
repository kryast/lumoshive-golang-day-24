package service

import "day-24/model"

func (bs *BookService) GetAllBooks() ([]model.Book, error) {
	return bs.RepoBook.GetAllBooks()
}
