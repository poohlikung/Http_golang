package main

import (
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
	"github.com/golang-jwt/jwt/v4"
)

// ประกาศ struct ในการกำหนดแม่แบบ
type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

var books []Book

func checkMiddleware(c *fiber.Ctx) error {
	// start := time.Now()
	// fmt.Printf("URL = %s,Method = %s,Time = %s\n", c.OriginalURL(), c.Method(), start)

	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	if claims["role"] != "admin" {
		return fiber.ErrUnauthorized
	}

	return c.Next()
}

func main() {
	app := fiber.New()

	books = append(books, Book{ID: 1, Title: "Poohlikung", Author: "Pooh"})

	books = append(books, Book{ID: 2, Title: "Paleerat", Author: "Nampung"})

	app.Post("/login", loginUser)

	// when login sucess middle are show

	// JWT Middleware
	app.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte(os.Getenv("JWT_SECRET")),
	}))

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

	// Create token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = user.Email
	claims["admin"] = "admin"
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	// Generate encoded token
	t, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{
		"message": "Login success",
		"token":   t,
	})
}
