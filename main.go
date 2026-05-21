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

	app.Get("/books", getBooks)

	app.Post("/books", createBooks)
	// one pice
	app.Get("/books/:id", getBook)
	// Put
	app.Put("/books/:id", updateBook)
	// Delete
	app.Delete("/books/:id", deleteBook)
	// Upload File
	app.Post("/upload", uploadFile)
	app.Listen(":8080")
}

func uploadFile(c *fiber.Ctx) error {
	file, err := c.FormFile("image")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())

	}
	err = c.SaveFile(file, "./uploads/"+file.Filename)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())

	}

	return c.SendString("Upload complete")
}
