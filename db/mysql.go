package db

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/mustafa-533/rest-api/model"
)

var ErrNotFound = errors.New("not found")

type MySQL struct {
	db *sqlx.DB
}

func LoadMySqlDB(url string) (*MySQL, error) {
	db, err := sqlx.Open("mysql", url)
	if err != nil {
		return nil, fmt.Errorf("error opening mysql connection: %w", err)
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)

	return &MySQL{
		db: db,
	}, nil
}

func (m *MySQL) GetAll() ([]model.Book, error) {
	var books []model.Book

	err := m.db.Select(&books, "SELECT * FROM books")
	if err != nil {
		return nil, err
	}

	return books, nil
}

func (m *MySQL) GetByID(id int) (*model.Book, error) {
	var book model.Book

	err := m.db.Get(&book, "SELECT * FROM books WHERE id = ?", id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrNotFound
		} else {
			return nil, err
		}
	}

	return &book, nil
}

func (m *MySQL) Create(book *model.Book) (*model.Book, error) {
	res, err := m.db.Exec("INSERT INTO books (title, author) VALUES (?, ?)", book.Title, book.Author)
	if err != nil {
		return nil, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}

	book.ID = int(id)

	return book, nil
}

func (m *MySQL) Update(book *model.Book, id int) error {
	_, err := m.db.Exec("UPDATE books SET title = ?, author = ? WHERE id = ?", book.Title, book.Author, id)
	if err != nil {
		return err
	}

	return nil
}

func (m *MySQL) Delete(id int) error {
	_, err := m.db.Exec("DELETE FROM books WHERE id = ?", id)
	if err != nil {
		return err
	}

	return nil
}
