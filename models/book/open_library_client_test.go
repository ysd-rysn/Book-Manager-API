package book

import "testing"

func TestOpenLibraryClient(t *testing.T) {
	isbn_10 := "0134190440"    // book title: "The Go Programming Language"
	isbn_13 := "9781491941195" // book title: "Concurrency in Go"
	if _, err := FetchBookData(isbn_10); err != nil {
		t.Fatal(err)
	}
	if _, err := FetchBookData(isbn_13); err != nil {
		t.Fatal(err)
	}
}
