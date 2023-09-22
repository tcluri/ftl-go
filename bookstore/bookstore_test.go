package bookstore_test

import (
	"bookstore"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestBook(t *testing.T) {
	t.Parallel()
	_ = bookstore.Book{
		Title:  "Spark Joy",
		Author: "Marie Kondo",
		Copies: 2,
	}
}

func TestBuy(t *testing.T) {
	t.Parallel()
	b := bookstore.Book{
		Title:  "Huckleberry Finn",
		Author: "Mark Twain",
		Copies: 5,
	}
	want := 4
	bought, err := bookstore.Buy(b)
	if err != nil {
		t.Fatal(err)
	}
	got := bought.Copies
	if want != got {
		t.Errorf("want %d copies after buying 1 copy from a stock of %d copies, got %d", want, b.Copies, got)
	}
}

func TestBuyErrorsIfNoCopiesLeft(t *testing.T) {
	t.Parallel()
	b := bookstore.Book{
		Title:  "Spark Joy",
		Author: "Marie Kondo",
		Copies: 0,
	}
	_, err := bookstore.Buy(b)
	if err == nil {
		t.Error("want error buying from zero copies, got nil")
	}
}

func TestGetAllBooks(t *testing.T) {
	t.Parallel()
	catalog := []bookstore.Book{
		{Title: "For the Love of Go"},
		{Title: "The Power of Go: Tools"},
	}
	want := []bookstore.Book{
		{Title: "For the Love of Go"},
		{Title: "The Power of Go: Tools"},
	}
	got := bookstore.GetAllBooks(catalog)
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestGetBook(t *testing.T) {
	t.Parallel()
	catalog := map[int]bookstore.Book{
		1: {
			ID:    1,
			Title: "For the Love of Go",
		},
		2: {
			ID:    2,
			Title: "The Power of Go: Tools",
		},
	}
	want := bookstore.Book{
		ID:    2,
		Title: "The Power of Go: Tools",
	}
	got := bookstore.GetBook(catalog, 2)
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}
