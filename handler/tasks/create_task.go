package tasks

import (
	"CA_Backend/database"
	"CA_Backend/models"
	"context"
	"log"
	"github.com/gofiber/fiber/v2"
)

func CreateTask(c *fiber.Ctx) error {
	task := new(models.Task)
	if err := c.BodyParser(task); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error":   err.Error(),
			"message": "Failed to parse JSON Body",
		})
	}
	ctx := context.Background()
	db, err := database.Connect()
	if err != nil {
		log.Fatal(err.Error())
		return c.Status(500).JSON(fiber.Map{"message": err.Error()})
	}
	collection := db.Collection("tasks")
	if res, err := collection.InsertOne(ctx, task); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error":   err.Error(),
			"message": "Failed to create task!!",
		})
	} else {
		return c.Status(201).JSON(fiber.Map{
			"id":      res.InsertedID,
			"message": "Task created successfully",
		})
	}
}
