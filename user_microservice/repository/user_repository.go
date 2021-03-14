package repository

import (
	"SA-final/user_microservice/models"
	"context"
	"errors"
	"github.com/jackc/pgx/v4/pgxpool"
)

type IUserRepository interface {
	RegisterUser(user *models.User) (*models.User, error)
	GetAllUsers() ([]*models.User, error)
	GetUserById(id int) (*models.User, error)
	GetUserByUsername(username string) (*models.User, error)
	FindByEmailAndPassword(email, password string) (*models.User, error)
	AddBook(user models.User, book models.Book) (*models.User, error)
}

type UserRepository struct {
	Pool *pgxpool.Pool
}

func (u *UserRepository) RegisterUser(user *models.User) (*models.User, error) {
	sql := `insert into users(email, username, password) values($1, $2, $3) returning id`
	res := u.Pool.QueryRow(context.Background(), sql, user.Email, user.Username, user.Password)
	var id int
	err := res.Scan(&id)
	if err != nil {
		return nil, err
	}
	user.ID = id
	user.BookIds = []int{}
	return user, err
}

func (u *UserRepository) GetUserById(id int) (*models.User, error) {
	panic("implement me")
}

func (u *UserRepository) GetUserByUsername(username string) (*models.User, error) {
	sql := `select * from users where username = $1 limit 1`
	user := &models.User{}
	res := u.Pool.QueryRow(context.Background(), sql, username)
	err := res.Scan(&user.ID, &user.Email, &user.Username, &user.Password)
	if err != nil {
		return nil, err
	}
	user, err = u.GetUserBooks(user)
	return user, nil
}

func (u *UserRepository) GetUserBooks(user *models.User) (*models.User, error) {
	sql := `select book_id from users_books where user_id = $1`
	rows, err := u.Pool.Query(context.Background(), sql, user.ID)
	defer rows.Close()
	for rows.Next() {
		var bookId int
		err = rows.Scan(&bookId)
		if err != nil {
			return nil, err
		}
		user.AddBook(bookId)
	}
	return user, err
}

func (u *UserRepository) FindByEmailAndPassword(email, password string) (*models.User, error) {
	panic("implement me")
}

func (u *UserRepository) AddBook(user models.User, book models.Book) (*models.User, error) {
	if user.IsContainsBook(&book) {
		return nil, errors.New("book already added")
	}
	user.AddBook(book.ID)
	sql := `insert into users_books(user_id, book_id) values($1, $2)`
	_ = u.Pool.QueryRow(context.Background(), sql, user.ID, book.ID)
	return &user, nil
}

func (u *UserRepository) GetAllUsers() ([]*models.User, error) {
	sql := `select * from users`
	rows, err := u.Pool.Query(context.Background(), sql)
	defer rows.Close()
	var users []*models.User
	for rows.Next() {
		user := &models.User{}
		err = rows.Scan(&user.ID, &user.Email, &user.Username, &user.Password)
		if err != nil {
			return nil, err
		}
		user, err = u.GetUserByUsername(user.Username)
		users = append(users, user)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

