package routes

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
)

func customLogger(ctx *fiber.Ctx) error {
	log.Println("This is example logging middleware")
	// Go to next middleware:
	return ctx.Next()
}

func SetRoutes(app *fiber.App) {
	api := app.Group("/api")

	api.Use("*", customLogger)

	api.Get("/metrics", monitor.New(monitor.Config{Title: "Metrics Page"}))

	api.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Welcome to the API",
		})
	})

	api.Get("/hello/:message", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Hello " + c.Params("message"),
		})
	})

}
