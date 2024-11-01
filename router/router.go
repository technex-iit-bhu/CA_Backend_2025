package router

import (
	"CA_Backend/handler"
	user_handler "CA_Backend/handler/users"
	"CA_Backend/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func Route(app *fiber.App) {
	app.Use(cors.New())
	app.Use("/api", middleware.Protected())
	api := app.Use("/api", logger.New())
	api.Get("/", handler.Hello)

	user := api.Group("/user")
	user.Get("/profile", user_handler.GetUserProfile)
	user.Post("/register", user_handler.CreateUser)
	user.Patch("/login", user_handler.LoginUser)
}
