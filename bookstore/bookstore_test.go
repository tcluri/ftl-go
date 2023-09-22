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
