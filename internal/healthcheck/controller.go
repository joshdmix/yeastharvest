package healthcheck

import "github.com/gofiber/fiber"

func Healthcheck(c *fiber.Ctx) {
	println("healthcheck")
}
