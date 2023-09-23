package bookstore

import (
	"errors"
	"fmt"
)

type Book struct {
	Title           string
	Author          string
	Copies          int
	ID              int
	PriceCents      int
	DiscountPercent int
}

func Buy(b Book) (Book, error) {
	if b.Copies == 0 {
		return Book{}, errors.New("no copies left")
	}
	b.Copies--
	return b, nil
}

func GetAllBooks(catalog map[int]Book) []Book {
	result := []Book{}
	for _, b := range catalog {
		result = append(result, b)
	}
	return result
}

func GetBook(catalog map[int]Book, Bookid int) (Book, error) {
	b, ok := catalog[Bookid]
	if !ok {
		return Book{}, fmt.Errorf("ID %d doesn't exist", Bookid)
	}
	return b, nil
}

func NetPriceCents(b Book) int {
	saving := b.PriceCents * b.DiscountPercent / 100
	return b.PriceCents - saving
}
