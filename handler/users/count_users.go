package user

import (
	"CA_Backend/database"
	"context"
	"github.com/gofiber/fiber/v2"
	"log"
	"go.mongodb.org/mongo-driver/bson"
)

func CountUsers(c *fiber.Ctx) error {
	ctx := context.Background()
	db, err := database.Connect()
	if err != nil {
		log.Fatal(err.Error())
		return c.Status(500).JSON(fiber.Map{
			"error":   err.Error(),
			"message": "Database connection error",
		})
	}
	collection := db.Collection("users")
	
	count, err := collection.CountDocuments(ctx, bson.D{})
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error":   err.Error(),
			"message": "Error counting users",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"count": count,
	})
}
