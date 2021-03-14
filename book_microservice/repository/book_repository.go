package repository

import (
	"SA-final/book_microservice/models"
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
)

type IBookRepository interface {
	GetAllBooks() ([]*models.Book, error)
	GetBookWithId(id int) (*models.Book, error)
	SearchBook(name string) ([]*models.Book, error)
	SearchBookWithName(name string) (*models.Book, error)
	CreateBook(book *models.Book) (*models.Book, error)
}

type BookRepository struct {
	Pool *pgxpool.Pool
}

func (b *BookRepository) GetAllBooks() ([]*models.Book, error) {
	sql := `select * from books`
	rows, err := b.Pool.Query(context.Background(), sql)
	defer rows.Close()
	var books []*models.Book
	for rows.Next() {
		book := &models.Book{}
		err = rows.Scan(&book.ID, &book.Name, &book.Author)
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

func (b *BookRepository) GetBookWithId(id int) (*models.Book, error) {
	sql := `select * from books where id = $1`
	book := &models.Book{}
	res := b.Pool.QueryRow(context.Background(), sql, id)
	err := res.Scan(&book.ID, &book.Name, &book.Author)

	if err != nil {
		return nil, err
	}
	return book, nil
}

func (b *BookRepository) SearchBook(name string) ([]*models.Book, error) {
	sql := `select * from books where name like '%` + name + `%'`
	rows, err := b.Pool.Query(context.Background(), sql)
	defer rows.Close()
	var books []*models.Book
	for rows.Next() {
		book := &models.Book{}
		err = rows.Scan(&book.ID, &book.Name, &book.Author)
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

func (b *BookRepository) SearchBookWithName(name string) (*models.Book, error) {
	sql := `select * from books where name like '%` + name + `%' limit 1`
	book := &models.Book{}
	res := b.Pool.QueryRow(context.Background(), sql)
	err := res.Scan(&book.ID, &book.Name, &book.Author)
	if err != nil {
		return nil, err
	}
	return book, nil
}

func (b *BookRepository) CreateBook(book *models.Book) (*models.Book, error) {
	sql := `insert into books(name, author) values($1, $2) returning id`
	res := b.Pool.QueryRow(context.Background(), sql, book.Name, book.Author)
	var id int
	err := res.Scan(&id)
	if err != nil {
		return nil, err
	}
	book.ID = id
	return book, nil
}
