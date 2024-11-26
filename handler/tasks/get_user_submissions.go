package tasks

import (
	"CA_Backend/database"
	"CA_Backend/models"
	"CA_Backend/utils"
	"context"
	"log"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetUserSubmissions(c *fiber.Ctx) error {
	ctx := context.Background()
	db, err := database.Connect()
	if err != nil {
		log.Fatal(err.Error())
		return c.Status(500).JSON(fiber.Map{
			"error":   err.Error(),
			"message": "Database connection error",
		})
	}

	tokenString := c.Get("Authorization")
	if len(tokenString) < 8 || tokenString[:7] != "Bearer " {
		return c.Status(400).JSON(fiber.Map{"message": "Authorization header missing or improperly formatted"})
	}
	token := tokenString[7:]

	username, _, _ := utils.DeserialiseUser(token)

	var user_submissions []models.TaskSubmission

	cursor, err := db.Collection("task_submissions").Find(ctx, bson.D{{Key: "User", Value: username}}, options.Find())
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error":   err.Error(),
			"message": "Failed to fetch submissions",
		})
	}
	defer cursor.Close(ctx)

	if err := cursor.All(ctx, &user_submissions); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error":   err.Error(),
			"message": "Error decoding submissions",
		})
	}

	return c.Status(200).JSON(user_submissions)
}
