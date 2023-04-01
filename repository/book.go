package repository

import (
	"eight-learn/model"
	"eight-learn/repository/query"
	"errors"
)

// clean architectures -> handler->service->repo

// interface Book
type BookRepo interface {
	GetBooks() ([]model.Book, error)
	CreateBook(in model.Book) (res model.Book, err error)
	GetBookById(id int64) (res model.Book, err error)
	UpdateBook(in model.Book) (res model.Book, err error)
	DeleteBook(id int64) (err error)
}

func (r Repo) GetBooks() ([]model.Book, error) {
	rows, err := r.db.Query(query.GetBooks)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	books := make([]model.Book, 0)
	for rows.Next() {
		var book model.Book
		err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Desc)
		if err != nil {
			return nil, err
		}
		books = append(books, book)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return books, nil
}

func (r Repo) CreateBook(in model.Book) (res model.Book, err error) {
	err = r.db.QueryRow(
		query.AddBook,
		in.Title,
		in.Author,
		in.Desc,
	).Scan(
		&res.ID,
		&res.Title,
		&res.Author,
		&res.Desc,
	)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (r Repo) GetBookById(id int64) (res model.Book, err error) {

	err = r.db.QueryRow(
		query.GetBookById,
		id,
	).Scan(
		&res.ID,
		&res.Title,
		&res.Author,
		&res.Desc,
	)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (r Repo) UpdateBook(in model.Book) (res model.Book, err error) {
	// Update book record in the database
	err = r.db.QueryRow(
		query.UpdateBook,
		in.Title,
		in.Author,
		in.Desc,
		in.ID,
	).Scan(
		&res.ID,
		&res.Title,
		&res.Author,
		&res.Desc,
	)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (r Repo) DeleteBook(id int64) (err error) {
	result, err := r.db.Exec(query.DeleteBook, id)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("no rows were affected")
	}
	return nil
}
