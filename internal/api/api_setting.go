package api

import (
	"oat431/try-go-discord-bot/internal/router"
	"os"

	"github.com/gofiber/fiber/v3"
)

func StartAPI() {
	app := fiber.New()

	router.SetupRouter(app)
	port := os.Getenv("API_PORT")
	error := app.Listen(":" + port)
	if error != nil {
		panic(error)
	}

}
