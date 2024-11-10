package tasks

import (
	"CA_Backend/database"
	"CA_Backend/models"
	"context"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func UpdateTask(c *fiber.Ctx) error {
	ctx := context.Background()
	db, err := database.Connect()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": err.Error()})
	}

	collection := db.Collection("tasks")
	updatedTask := new(models.User)
	if err := c.BodyParser(updatedTask); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error":   err.Error(),
			"message": "Failed to parse JSON Body",
		})
	}

	task := new(models.Task)
	task_id:=c.Params("task_id")
	objectId, _ := primitive.ObjectIDFromHex(task_id)
	if err := collection.FindOne(ctx, bson.D{{Key: "_id", Value: objectId}}).Decode(task); err != nil {
		return c.Status(404).JSON(fiber.Map{"message": "Task not found"})
	}

	if _, err := collection.UpdateOne(ctx, bson.D{{Key: "_id", Value: objectId}}, updatedTask); err != nil {
		return c.Status(500).JSON(fiber.Map{"message": err.Error()})
	} else {
		return c.Status(200).JSON(fiber.Map{
			"message": "Task updated Successfully",
			"task":    updatedTask,
		})
	}
}
