package user

import (
	"CA_Backend/database"
	"CA_Backend/models"
	"CA_Backend/utils"
	"context"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func LoginUser(c *fiber.Ctx) error {
	db, err := database.Connect()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": err.Error()})
	}

	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error":   err.Error(),
			"message": "Failed to parse JSON Body",
		})
	}

	var result models.User
	err = db.Collection("users").FindOne(context.Background(), bson.D{{Key: "username", Value: user.Username}}).Decode(&result)

	if err != nil {
		return c.Status(404).JSON(fiber.Map{"message": "invalid username"})
	}
	if !utils.CheckPassword(user.Password, result.Password) {
		return c.Status(404).JSON(fiber.Map{"message": "invalid password"})
	} else {
		token, err := utils.SerialiseUser(result.Username)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"message": err.Error()})
		}
		return c.Status(200).JSON(fiber.Map{
			"token":   token,
			"message": "Login successful",
		})
	}
}
