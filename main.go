package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type book struct {
	ID       string `json:"id,omitempty"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Quantity int    `json:"quantity"`
}

var books = []book{
	{
		ID:       "1",
		Title:    "In Search of lost time",
		Author:   "Marcel Proust",
		Quantity: 2,
	},
	{
		ID:       "2",
		Title:    "In Search of lost time",
		Author:   "Marcel Proust",
		Quantity: 2,
	},
	{
		ID:       "3",
		Title:    "War and Peace",
		Author:   "Cemo Ben",
		Quantity: 6,
	},
}

func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

func main() {
	router := gin.Default()
	router.GET("/books", getBooks)
	router.Run("localhost:8080")
}
