package service

import (
	"day-24/model"
	"fmt"
)

func (bs *BookService) GetAllBooks() ([]model.Book, error) {
	return bs.RepoBook.GetAllBooks()
}

func (bs *BookService) GetBookByID(id int) (*model.Book, error) {
	book, err := bs.RepoBook.FindByID(id)
	if err != nil {
		return nil, fmt.Errorf("could not find book with id %d: %v", id, err)
	}
	return book, nil
}
