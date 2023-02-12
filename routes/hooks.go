package routes

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func SetHooks(app *fiber.App) {
	app.Hooks().OnListen(func() error {
		log.Print("Server has started")
		return nil
	})

}
