package main

import (
	"log"

	"github.com/gofiber/fiber"
	"github.com/gofiber/fiber/middleware"
	_ "github.com/lib/pq"

	"github.com/joshdmix/harvestyeast/database"
	"github.com/joshdmix/harvestyeast/router"
)

func main() {
	if err := database.Connect(); err != nil {
		log.Fatal(err)
	}

	app := fiber.New()

	app.Use(middleware.Logger())

	router.SetupRoutes(app)

	app.Listen(3000)
}

func handler() func(*fiber.Ctx) {
	return func(c *fiber.Ctx) {

	}
}
