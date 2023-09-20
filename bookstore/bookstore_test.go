package bookstore_test

import (
	"bookstore"
	"testing"
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
	bought := bookstore.Buy(b)
	got := bought.Copies
	if want != got {
		t.Errorf("want %d copies after buying 1 copy from a stock of %d copies, got %d", want, b.Copies, got)
	}
}
