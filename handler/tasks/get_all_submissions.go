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

func GetAllSubmissions(c *fiber.Ctx) error {
	ctx := context.Background()
	db, err := database.Connect()
	if err != nil {
		log.Fatal(err.Error())
		return c.Status(500).JSON(fiber.Map{
			"error":   err.Error(),
			"message": "Database connection error",
		})
	}

	var all_submissions []models.TaskSubmission

	cursor, err := db.Collection("task_submissions").Find(ctx, bson.D{}, options.Find())
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error":   err.Error(),
			"message": "Failed to fetch submissions",
		})
	}
	defer cursor.Close(ctx)

	if err := cursor.All(ctx, &all_submissions); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error":   err.Error(),
			"message": "Error decoding submissions",
		})
	}

	return c.Status(200).JSON(all_submissions)
}
