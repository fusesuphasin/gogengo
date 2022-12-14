package routes

import (
	"context"
	"gogenerate/controller"
	"net/http"

	//swagger "github.com/arsmn/fiber-swagger/v2"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoute(app *fiber.App, ctx context.Context) {
	GenerateStructureController := controller.NewGenerateStructureController(app)
	GenerateStructureController.GenerateStructureRouter()

	app.Get("*", func(c *fiber.Ctx) error {
		return c.Status(http.StatusOK).JSON(struct{ Status string }{"Go generate is working"})
	})
	app.Get("/health", controller.HealthCheckHanlder)
}  