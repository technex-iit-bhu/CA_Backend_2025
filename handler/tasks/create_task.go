package tasks

import (
	"CA_Portal_backend/database"
	"CA_Portal_backend/models"
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
)

func CreateTask(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	db, err := database.Connect()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error":   err.Error(),
			"message": "Database connection failed",
		})
	}

	task := new(models.Task)
	if err := c.BodyParser(task); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error":   err.Error(),
			"message": "Failed to parse request body",
		})
	}
	task.Status = "active"
	collection := db.Collection("tasks")

	result, err := collection.InsertOne(ctx, task)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error":   err.Error(),
			"message": "Failed to create task",
		})
	}

	return c.Status(201).JSON(fiber.Map{
		"id":      result.InsertedID,
		"message": "Task created successfully",
	})
}
