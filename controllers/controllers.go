package controllers

import (
	"book-manager-api/models"
	modelsBook "book-manager-api/models/book"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateBook(c *gin.Context) {
	var book modelsBook.Book

	err := c.BindJSON(&book)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message":  "Book creation failed!",
			"required": "isbn, type, own",
		})
		return
	}

	// fetch book data from open library api
	bookData, err := modelsBook.FetchBookData(book.ISBN)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Book creation failed!",
		})
	}
	book.BindBookInfo(bookData)

	models.DB.Create(&book)

	c.JSON(http.StatusOK, gin.H{
		"book": book,
	})
}

func GetBooks(c *gin.Context) {
	var books []modelsBook.Book
	models.DB.Find(&books)

	c.JSON(http.StatusOK, gin.H{
		"books": books,
	})

}

func GetBook(c *gin.Context) {
	var book modelsBook.Book

	if err := models.DB.Where("isbn = ?", c.Param("isbn")).First(&book).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "No book found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "book details by isbn",
		"book":    book,
	})
}

func UpdateBook(c *gin.Context) {
	var book modelsBook.Book

	if err := models.DB.Where("isbn = ?", c.Param("isbn")).First(&book).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "No book found",
		})
		return
	}

	err := c.BindJSON(&book)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message":  "book update failed!",
			"required": "isbn, type, own",
		})
		return
	}

	models.DB.Model(&book).Updates(book)

	c.JSON(http.StatusOK, gin.H{
		"message": "book successfully updated!",
		"book":    book,
	})
}

func DeleteBook(c *gin.Context) {
	var book modelsBook.Book

	if err := models.DB.Where("isbn = ?", c.Param("isbn")).First(&book).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "No book found",
		})
		return
	}

	models.DB.Delete(&book)

	c.JSON(http.StatusOK, gin.H{
		"message": "book successfully deleted!",
	})
}
