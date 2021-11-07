package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func initAPI(c *fiber.Ctx) error {
	return c.SendString("Restful API using GO")
}

func main() {
	app := fiber.New()

	app.Get("/api", initAPI)

	log.Fatal(app.Listen(":3000"))
}
