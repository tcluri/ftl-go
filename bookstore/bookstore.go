package bookstore

import (
	"errors"
)

type Book struct {
	Title  string
	Author string
	Copies int
	ID     int
}

func Buy(b Book) (Book, error) {
	if b.Copies == 0 {
		return Book{}, errors.New("no copies left")
	}
	b.Copies--
	return b, nil
}

func GetAllBooks(catalog []Book) []Book {
	return catalog
}

func GetBook(catalog []Book, bookid int) Book {
	for _, book := range catalog {
		if book.ID == bookid {
			return book
		}
	}
	return Book{}
}
