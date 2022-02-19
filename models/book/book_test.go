package book

import (
	"fmt"
	"testing"
)

func TestBindBookInfo(t *testing.T) {
	isbn_10 := "0134190440" // book title: "The Go Programming Language"

	var book Book

	if bookData, err := FetchBookData(isbn_10); err != nil {
		t.Fatal(err)
	} else {
		book.BindBookInfo(bookData)
		fmt.Printf("Title: %s, CoverURL: %s, NumPages: %d\n", book.Title, book.CoverURL, book.NumPages)

		is_zero_value := (book.Title == "") || (book.CoverURL == "") || (book.NumPages == 0)
		if is_zero_value {
			t.Fatal("any of them is zero value")
		}
	}
}
