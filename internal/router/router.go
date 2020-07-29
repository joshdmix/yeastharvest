package router

import (
	"github.com/gofiber/fiber"

	"github.com/joshdmix/yeastharvest/internal/healthcheck"
)

//SetupRoutes func
func SetupRoutes(app *fiber.App) {
	app.Get("/healthcheck", healthcheck.Healthcheck)
}
