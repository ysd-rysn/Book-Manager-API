package main

import (
	"book-manager-api/controllers"
	"book-manager-api/models"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	models.ConnectDB()

	r := gin.Default()

	r.POST("/books", controllers.CreateBook)
	r.GET("/books", controllers.GetBooks)
	r.GET("/books/:isbn", controllers.GetBook)
	r.PATCH("/books/:isbn", controllers.UpdateBook)
	r.DELETE("/books/:isbn", controllers.DeleteBook)

	if os.Getenv("PORT") != "" {
		r.Run(":" + os.Getenv("PORT"))
	} else {
		r.Run()
	}
}
