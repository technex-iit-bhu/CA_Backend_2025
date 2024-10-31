package router

import (
	"CA_Backend/handler"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	user_handler "CA_Backend/handler/users"
)

func Route(app *fiber.App) {
	app.Use(cors.New())

	api := app.Use("/api", logger.New())
	api.Get("/", handler.Hello)

	user := api.Group("/user")
	user.Post("/register", user_handler.CreateUser)
	user.Get("/profile", user_handler.GetUserFromToken)
}
