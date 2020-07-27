package router

import (
	"github.com/gofiber/fiber"
	"github.com/gofiber/logger"

	"github.com/joshdmix/harvestyeast/handler"
	"github.com/joshdmix/harvestyeast/middleware"
)

//SetupRoutes func
func SetupRoutes(app *fiber.App) {
	// Middleware
	api := app.Group("/api", logger.New(), middleware.AuthReq())

	// routes
	api.Get("/", handler.GetAllProducts)
	api.Get("/:id", handler.GetSingleProduct)
	api.Post("/", handler.CreateProduct)
	api.Delete("/:id", handler.DeleteProduct)

}
