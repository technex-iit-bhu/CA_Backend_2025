package router

import (
	"CA_Backend/handler"
	task_handler "CA_Backend/handler/tasks"
	user_handler "CA_Backend/handler/users"
	"CA_Backend/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func Route(app *fiber.App) {
	app.Use(cors.New())
	api := app.Group("/api", logger.New())
	api.Get("/", handler.Hello)

	user := api.Group("/user")
	user.Use("/user", middleware.Protected())
	user.Get("/profile", user_handler.GetUserProfile)
	user.Post("/register", user_handler.CreateUser)
	user.Patch("/login", user_handler.LoginUser)
	user.Patch("/setReferral", user_handler.SetReferral)

	password := user.Group("/password")
	password.Post("/recovery", user_handler.RequestPasswordRecovery)
	password.Post("/reset", user_handler.ResetPassword)

	tasks := api.Group("/tasks")
	tasks.Get("/", task_handler.GetAllTasks)
}
