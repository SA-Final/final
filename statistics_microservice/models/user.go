package models

import (
	"fmt"
)

type User struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
	BookIds  []int  `json:"bookIds"`
}

func (u *User) AddBook(id int) {
	u.BookIds = append(u.BookIds, id)
}

func (u *User) IsContainsBook(checkingBook *Book) bool {
	for _, id := range u.BookIds {
		if id == checkingBook.ID {
			fmt.Print("true")
			return true
		}
	}
	return false
}
