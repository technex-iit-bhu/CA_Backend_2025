package tasks

import (
	"CA_Backend/database"
	"CA_Backend/models"
	"context"
	"log"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetTask(c *fiber.Ctx) error {
	ctx := context.Background()
	task_id:=c.Params("task_id")
	objectId, _ := primitive.ObjectIDFromHex(task_id)	

	db, err := database.Connect()
	if err != nil {
		log.Fatal(err.Error())
		return c.Status(500).JSON(fiber.Map{
			"error":   err.Error(),
			"message": "Database connection error",
		})
	}

	var result models.Task
	if err := db.Collection("tasks").FindOne(ctx, bson.D{{Key: "_id", Value: objectId}}).Decode(&result); err != nil {
		return c.Status(404).JSON(fiber.Map{"message": "Task does not exist"})
	}
	return c.Status(200).JSON(fiber.Map{"data": result})
}
