package main

import (
	"github.com/gofiber/fiber/v2"
)

// ประกาศ struct ในการกำหนดแม่แบบ
type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

var books []Book

func main() {
	app := fiber.New()

	books = append(books, Book{ID: 1, Title: "Poohlikung", Author: "Pooh"})

	books = append(books, Book{ID: 2, Title: "Paleerat", Author: "Nampung"})

	app.Get("/books", func(c *fiber.Ctx) error {
		return c.JSON(books)
	})

	app.Listen(":8080")
}
