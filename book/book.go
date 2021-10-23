package book 

import "github.com/gofiber/fiber/v2"

func GetBooks(c *fiber.Ctx){
	c.Send("All books")
}
func GetBook(c *fiber.Ctx) {
	c.Send("All a books")
}
func NewBook(c *fiber.Ctx) {
	c.Send("Add new books")
}
func DeleteBook(c *fiber.Ctx) {
	c.Send("Delete book")
}
