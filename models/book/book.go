package book

import (
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	ISBN         string `json:"isbn" binding:"required"`
	Title        string `json:"title"`
	CoverURL     string `json:"cover_url"`
	NumPages     uint   `json:"number_of_pages"`
	Type         string `json:"type" binding:"required"` // value is "printed book" or "ebook"
	Format       string `json:"format"`                  // set if type is ebook. value is "pdf" or "epub"
	Tag          string `json:"tag"`
	MarkdownMemo string `json:"markdown_memo"`
	Own          bool   `json:"own" binding:"required"` // whether own the book
}

// bind a book info from open library
func (b *Book) BindBookInfo(bookData map[string]interface{}) {
	b.Title = bookData["title"].(string)
	b.CoverURL = bookData["cover"].(map[string]interface{})["large"].(string)
	b.NumPages = uint(bookData["number_of_pages"].(float64))
}
