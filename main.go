package main

import (
	"github.com/Yohenba18/go-project.git/book"
	"github.com/gofiber/fiber/v2"
)

func hello(c *fiber.Ctx)error{
	return c.SendString("Hello, World!")
}

func setupRoutes(app *fiber.App) {
  app.Get("/api/v1/book", book.GetBooks)
  app.Get("/api/v1/book/:id", book.GetBook)
  app.Post("/api/v1/book/", book.NewBook)
  app.Delete("/api/v1/book/:id", book.DeleleBook)

}

func main() {
  app := fiber.New()

  setupRoutes(app)

  app.Listen(":3000")
}