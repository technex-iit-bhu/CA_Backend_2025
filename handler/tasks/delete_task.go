package tasks

import (
	"CA_Backend/database"
	"context"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func DeleteTask(c *fiber.Ctx) error {
	ctx := context.Background()
	db, err := database.Connect()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": err.Error()})
	}

	collection := db.Collection("tasks")
	task_id := c.Params("task_id")
	objectId, _ := primitive.ObjectIDFromHex(task_id)

	result, err := collection.DeleteOne(ctx, bson.D{{Key: "_id", Value: objectId}})
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": err.Error()})
	}

	if result.DeletedCount == 0 {
		return c.Status(404).JSON(fiber.Map{"message": "Task not found"})
	}

	return c.Status(200).JSON(fiber.Map{"message": "Task deleted successfully"})
}
