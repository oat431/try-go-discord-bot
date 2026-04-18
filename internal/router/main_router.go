package router

import "github.com/gofiber/fiber/v3"

func SetupRouter(app *fiber.App) {
	api := app.Group("/discord")
	hook := api.Group("/hook")

	RegisterHealthCheckRoute(hook)
	RegisterSendMessageRoute(hook)

}
