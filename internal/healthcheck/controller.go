package healthcheck

import "github.com/gofiber/fiber"

// Healthcheck function
func Healthcheck(c *fiber.Ctx) {
	println("healthcheck")
}
