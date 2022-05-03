package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/leksyking/go-fiber-mongo-hrms/database"
	"github.com/leksyking/go-fiber-mongo-hrms/routes"
)

func main() {
	if err := database.Connect(); err != nil {
		log.Fatal(err)
	}
	app := fiber.New()
	routes.EmployeeRoutes(app)
	app.Use("/", func(c *fiber.Ctx) error {
		return c.SendString("Welcome to fiber HRMS")
	})
	log.Fatal(app.Listen(":3000"))
}
