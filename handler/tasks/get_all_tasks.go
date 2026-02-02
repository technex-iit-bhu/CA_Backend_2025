package tasks

import (
	"CA_Portal_backend/database"
	"CA_Portal_backend/models"
	"context"
	"log"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetAllTasks(c *fiber.Ctx) error {
	ctx := context.Background()
	db, err := database.Connect()
	if err != nil {
		log.Fatal(err.Error())
		return c.Status(500).JSON(fiber.Map{
			"error":   err.Error(),
			"message": "Database connection error",
		})
	}

	var tasks []models.Task

	cursor, err := db.Collection("tasks").Find(ctx, bson.M{}, options.Find())
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error":   err.Error(),
			"message": "Failed to fetch tasks",
		})
	}
	defer cursor.Close(ctx)

	if err := cursor.All(ctx, &tasks); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error":   err.Error(),
			"message": "Error decoding tasks",
		})
	}

	return c.Status(200).JSON(tasks)
}
