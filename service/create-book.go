package service

import (
	"day-24/model"
	"day-24/repository"
	"fmt"
)

type BookService struct {
	RepoBook repository.BookRepositoryDB
}

func NewBookService(repo repository.BookRepositoryDB) BookService {
	return BookService{RepoBook: repo}
}

func (bs *BookService) CreateBook(book model.Book) error {
	err := bs.RepoBook.CreateDataBook(book)
	if err != nil {
		return fmt.Errorf("error creating book: %v", err)
	}
	return nil
}
