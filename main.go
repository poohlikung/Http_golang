package main

import (
	"fmt"
	"time"
  "github.com/gofiber/jwt/v2"
  "github.com/golang-jwt/jwt/v4"
	"github.com/gofiber/fiber/v2"
	"os"
)

// ประกาศ struct ในการกำหนดแม่แบบ
type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

var books []Book

func checkMiddleware(c *fiber.Ctx) error {
	start := time.Now()
	fmt.Printf("URL = %s,Method = %s,Time = %s\n", c.OriginalURL(), c.Method(), start)
	return c.Next()
}

func main() {
	app := fiber.New()

	books = append(books, Book{ID: 1, Title: "Poohlikung", Author: "Pooh"})

	books = append(books, Book{ID: 2, Title: "Paleerat", Author: "Nampung"})

	app.Post("/login", loginUser)
	  // JWT Middleware
  app.Use(jwtware.New(jwtware.Config{
    SigningKey: []byte(os.Getenv("JWT_SECRET")),
  }))
	// when login sucess middle are show
	app.Use(checkMiddleware)
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

type User struct {
	Email    string `json:"Email"`
	Password string `json:"Password"`
}

var memberUser = User{
	Email:    "sorawit@gmail.com",
	Password: "123456",
}

func loginUser(c *fiber.Ctx) error {
	user := new(User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	if user.Email != memberUser.Email || user.Password != memberUser.Password {
		return fiber.ErrUnauthorized
	}
	return c.JSON(fiber.Map{
		"message": "login sucess",
	})
}
