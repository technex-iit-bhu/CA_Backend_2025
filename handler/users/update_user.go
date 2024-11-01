package user

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"CA_Backend/database"
	"CA_Backend/models"
)

func UpdateUserDetails(c *fiber.Ctx) error {
	ctx := context.Background()
	db, err := database.Connect()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": err.Error()})
	}

	collection := db.Collection("users")
	updatedUser := new(models.User)
	if err := c.BodyParser(updatedUser); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error":   err.Error(),
			"message": "Failed to parse JSON Body",
		})
	}
	
	user := new(models.User)
	if err := collection.FindOne(ctx, models.User{Username: updatedUser.Username}).Decode(user); err != nil {
		return c.Status(404).JSON(fiber.Map{"message": "User not found"})
	}
	
	if _, err := collection.UpdateOne(ctx, models.User{Username: updatedUser.Username}, updatedUser); err != nil {
		return c.Status(500).JSON(fiber.Map{"message": err.Error()})
	} else {
		return c.Status(200).JSON(fiber.Map{
			"message": "User updated Successfully",
			"user":    updatedUser,
		})
	}
}