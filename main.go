package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/riju-stone/go-rest/database"
)

func initAPI(c *fiber.Ctx) error {
	return c.SendString("Restful API using GO")
}

func main() {
	database.ConnectDB()
	app := fiber.New()

	app.Get("/api", initAPI)

	log.Fatal(app.Listen(":3000"))
}
