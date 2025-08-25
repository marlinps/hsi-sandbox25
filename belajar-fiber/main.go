package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {
	// TODO: inisialisasi framework Fiber
	app := fiber.New()

	// TODO: routing
	// URL http://localhost:3000/api/
	// Return "Ahlan Wa Sahlan"
	app.Get("/api/", func(c *fiber.Ctx) error {
		return c.SendString("Ahlan Wa Sahlan")
	})

	// Parameter URL didefinisikan (:)
	app.Get("/api/:p1", func(c *fiber.Ctx) error {
		// TODO: untuk mendapatkan value dari parameter
		getParam := c.Params("p1")
		return c.SendString("Parameter yang dikirimkan adalah : " + getParam)
	})

	// Reserved PORT
	app.Listen(":3000")
}
