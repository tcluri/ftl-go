package bookstore

import (
	"errors"
	"fmt"
)

type Category int

const (
	CategoryAutobiography Category = iota
	CategoryLargePrintRomance
	CategoryParticlePhysics
)

var validCategory = map[Category]bool{
	CategoryAutobiography:     true,
	CategoryLargePrintRomance: true,
	CategoryParticlePhysics:   true,
}

type Book struct {
	Title           string
	Author          string
	Copies          int
	ID              int
	PriceCents      int
	DiscountPercent int
	category        Category
}

type Catalog map[int]Book

func Buy(b Book) (Book, error) {
	if b.Copies == 0 {
		return Book{}, errors.New("no copies left")
	}
	b.Copies--
	return b, nil
}

func (catalog Catalog) GetAllBooks() []Book {
	result := []Book{}
	for _, b := range catalog {
		result = append(result, b)
	}
	return result
}

func (catalog Catalog) GetBook(Bookid int) (Book, error) {
	b, ok := catalog[Bookid]
	if !ok {
		return Book{}, fmt.Errorf("ID %d doesn't exist", Bookid)
	}
	return b, nil
}

func (b Book) NetPriceCents() int {
	saving := b.PriceCents * b.DiscountPercent / 100
	return b.PriceCents - saving
}

func (b *Book) SetPriceCents(price int) error {
	if price < 0 {
		return fmt.Errorf("price in cents cannot be less than zero, got %d", price)
	}
	b.PriceCents = price
	return nil
}

func (b Book) Category() Category {
	return b.category
}

func (b *Book) SetCategory(category Category) error {
	if !validCategory[category] {
		return fmt.Errorf("unknown category %v", category)
	}
	b.category = category
	return nil
}
