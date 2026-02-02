package router

import (
	"CA_Portal_backend/handler"
	task_handler "CA_Portal_backend/handler/tasks"
	user_handler "CA_Portal_backend/handler/users"
	"CA_Portal_backend/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func Route(app *fiber.App) {
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
		AllowMethods: "GET, POST, PATCH, DELETE",
	}))
	api := app.Group("/api", logger.New())
	api.Get("/", handler.Hello)
	api.Get("/leaderboard", handler.GetLeaderboard)

	user := api.Group("/user")
	user.Use("/user", middleware.Protected())
	user.Get("/profile", user_handler.GetUserProfile)
	user.Post("/register", user_handler.CreateUser)
	user.Patch("/login", user_handler.LoginUser)
	user.Patch("/setReferral", user_handler.SetReferral)
	user.Patch("/update", user_handler.UpdateUserDetails)
	user.Get("/count", user_handler.CountUsers)

	password := user.Group("/password")
	password.Post("/recovery", user_handler.RequestPasswordRecovery)
	password.Post("/reset", user_handler.ResetPassword)

	tasks := api.Group("/tasks")
	tasks.Get("/", task_handler.GetAllTasks)
	tasks.Post("/create", task_handler.CreateTask)
	tasks.Get("/task/:task_id", task_handler.GetTask)
	tasks.Post("/update/:task_id", task_handler.UpdateTask)
	tasks.Delete("/task/:task_id", task_handler.DeleteTask)

	submissions := api.Group("/submissions")
	submissions.Post("/submit", task_handler.SubmitTask)
	submissions.Get("/get_user_submissions", task_handler.GetUserSubmissions)
	submissions.Get("/verify/:submission_id", task_handler.VerifySubmission)
	submissions.Get("/all", task_handler.GetAllSubmissions)
	submissions.Post("/comment/:submission_id", task_handler.AddAdminComment) // NEW: Add this line
}
