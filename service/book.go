package service

// usecase

import "eight-learn/model"

type BookService interface {
	GetBooks() ([]model.Book, error)
	CreateBook(in model.Book) (res model.Book, err error)
	GetBookById(id int64) (model.Book, error)
	UpdateBook(id int64, in model.Book) (res model.Book, err error)
	DeleteBook(id int64) (err error)
}

func (s *Service) CreateBook(in model.Book) (res model.Book, err error) {
	return s.repo.CreateBook(in)
}

func (s *Service) GetBookById(id int64) (res model.Book, err error) {
	return s.repo.GetBookById(id)
}

func (s Service) GetBooks() ([]model.Book, error) {
	return s.repo.GetBooks()
}

func (s *Service) UpdateBook(id int64, in model.Book) (res model.Book, err error) {
	return s.repo.UpdateBook(id, in)
}

func (s *Service) DeleteBook(id int64) (err error) {
	return s.repo.DeleteBook(id)
}
