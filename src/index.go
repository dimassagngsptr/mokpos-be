package src

import (
	"transaction-be/src/configs"
	"transaction-be/src/helpers"
	"transaction-be/src/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/helmet"
)

func App() *fiber.App {
	app := fiber.New()

	app.Use(helmet.New())

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
	}))

	configs.InitDB()
	helpers.Migration()
	routes.SetupRoutes(app)

	return app
}
