package handlers

import "github.com/gofiber/fiber/v2"

func GetHealth(c *fiber.Ctx) (err error) {
	return c.JSON(fiber.Map{
		"status":  "200",
		"message": "up and running",
	})
}
