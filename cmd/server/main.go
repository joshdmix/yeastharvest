package main

import (
	"log"

	"github.com/gofiber/fiber"

	"github.com/joshdmix/yeastharvest/internal/router"
)

func main() {
	app := fiber.New()
	router.SetupRoutes(app)
	if err := app.Listen(8000); err != nil {
		log.Fatal(err)
	}
}
