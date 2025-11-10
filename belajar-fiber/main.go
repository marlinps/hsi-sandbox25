package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// TODO: inisialisasi framework Fiber
	app := fiber.New()

	// TODO: Middleware (fungsi: Untuk Log Request disini, selain untuk autentifikasi/autorisasi dan sebagainya)
	app.Use(func(c *fiber.Ctx) error {
		log.Print("Middleware Berjalan Sebelum ke Routing")

		return c.Next() // TODO: selalu ditulis diakhir, berfungsi untuk meneruskan request ke routing yang dituju
	})

	// TODO: routing
	// URL http://localhost:3000/api/
	// Return "Ahlan Wa Sahlan"
	app.Get("/api/", func(c *fiber.Ctx) error {
		log.Print("Routing ke URL /api/") // TODO: mengetest middleware
		return c.SendString("Ahlan Wa Sahlan")
	})

	// Parameter URL Sample 1 didefinisikan (:p1)
	app.Get("/api/:p1", func(c *fiber.Ctx) error {
		// TODO: untuk mendapatkan value dari parameter
		getParam := c.Params("p1")
		return c.SendString("Parameter yang dikirimkan adalah : " + getParam)
	})

	// Parameter URL Sample 1 (Multiple Parameter URL)
	app.Get("/api/:categoryId/:productId", func(c *fiber.Ctx) error {
		getParam1 := c.Params("categoryId")
		getParam2 := c.Params("productId")
		return c.SendString("Paramater yang dikirimkan adalah Category ID: " + getParam1 + ", Product ID: " + getParam2)
	})

	// Request:

	//gitGET http://localhost:3000/api

	// Response:

	// TODO: Parameter URL Sample 3 (Optional Parameter) -> :id (TODO: ?) parameter in the route is optional.
	app.Get("/api2/opsional/:opsionalData?", func(c *fiber.Ctx) error {
		getParam := c.Params("opsionalData", "Default  Data")
		return c.SendString("Parameter yang dikirimkan adalah : " + getParam)

	})

	// TODO: biasa bodynya berbentuk struct
	type BodySample struct {
		Name    string `json:"name"` // TODO: auto mapping json key "name" ke struct Name
		Address string `json:"address"`
	}

	// Handler Post
	app.Post("/api/info", func(c *fiber.Ctx) error {
		// TODO: Inisialisasi object baru
		BodySample := new(BodySample)
		if err := c.BodyParser(BodySample); err != nil {
			return c.Status(fiber.StatusBadRequest).SendString(err.Error())
		}

		// return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		return c.JSON(fiber.Map{
			"status":  "success",
			"message": "Data berhasil diterima",
			"name":    BodySample.Name,
			"address": BodySample.Address,
		})
	})

	// Reserved PORT
	app.Listen(":3000")
}
