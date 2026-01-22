package user

import (
	"CA_Backend/database"
	"CA_Backend/models"
	"CA_Backend/utils"
	"context"
	"log"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func GetUserProfile(c *fiber.Ctx) error {
	ctx := context.Background()
	tokenString := c.Get("Authorization")
	log.Printf("Authorization header: %s\n", tokenString)
	if len(tokenString) < 8 || tokenString[:7] != "Bearer " {
		return c.Status(400).JSON(fiber.Map{"message": "Authorization header missing or improperly formatted"})
	}
	token := tokenString[7:]

	username, err := utils.DeserialiseUser(token)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "Invalid or expired token",
			"error": err.Error(),
		})
	}

	db, err := database.Connect()
	if err != nil {
		log.Fatal(err.Error())
		return c.Status(500).JSON(fiber.Map{
			"error":   err.Error(),
			"message": "Database connection error",
		})
	}

	// fetch user
	var user models.User
	if err := db.Collection("users").FindOne(
		ctx,
		bson.D{{Key: "username", Value: username}},
	).Decode(&user); err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "User does not exist",
		})
	}

	// calculate rank
	count, err := db.Collection("users").CountDocuments(
		ctx,
		bson.D{{Key: "points", Value: bson.D{{Key: "$gt", Value: user.Points}}}},
	)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Failed to calculate rank",
		})
	}

	rank := count + 1

	var result models.User
	if err := db.Collection("users").FindOne(ctx, bson.D{{Key: "username", Value: username}}).Decode(&result); err != nil {
		return c.Status(404).JSON(fiber.Map{"message": "User does not exist"})
	}
	return c.Status(200).JSON(fiber.Map{
		"user": fiber.Map{
			"id":       result.ID,
			"username": result.Username,
			"name":     result.Name,
			"email":    result.Email,
			"points":   result.Points,
			"rank":     rank,
			"ca_id":    result.CA_ID,
			"college":  result.Institute,
			"phone":    result.PhoneNumber,
			"joinedAt": result.CreatedAt,
		},
	})
}
