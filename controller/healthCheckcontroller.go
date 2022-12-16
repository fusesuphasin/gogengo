package controller

import (
	"net/http"

	fiber "github.com/gofiber/fiber/v2"
)

func HealthCheckHanlder(c *fiber.Ctx) error {
	return c.Status(http.StatusOK).JSON(struct{ Status string }{"PromptPay is working"})
}
