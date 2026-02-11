package user

import (
	"CA_Portal_backend/database"
	"CA_Portal_backend/models"
	"CA_Portal_backend/utils"
	"context"
	// "log"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func RequestPasswordRecovery(c *fiber.Ctx) error {
	type request struct {
		Email string `json:"email"`
	}

	var req request
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "Invalid request format"})
	}

	db, err := database.Connect()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Database connection error"})
	}

	var user models.User
	err = db.Collection("users").FindOne(context.Background(), bson.D{{Key: "email", Value: req.Email}}).Decode(&user)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"message": "Email not registered"})
	}
	// log.Printf("User found for email %s: %s", req.Email, user.Username)

	token, err := utils.SerialiseRecovery(user.Username)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Failed to generate recovery token"})
	}

	if err = utils.RecoveryMail(req.Email, token); err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Failed to send recovery email"})
	}

	return c.Status(200).JSON(fiber.Map{"message": "Recovery email sent successfully"})
}

func ResetPassword(c *fiber.Ctx) error {
	type resetRequest struct {
		Token       string `json:"token"`
		NewPassword string `json:"new_password"`
	}

	var req resetRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "Invalid request format"})
	}

	username, err := utils.DeserialiseRecovery(req.Token)
	if err != nil {
		return c.Status(401).JSON(fiber.Map{"message": "Invalid or expired token"})
	}

	if !utils.IsSafe(req.NewPassword) {
		return c.Status(400).JSON(fiber.Map{
			"message": "Password must contain 8 characters, 1 uppercase & 1 lowercase letter, 1 number and 1 special character",
		})
	}

	db, err := database.Connect()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Database connection error"})
	}

	hashedPassword := utils.HashPassword(req.NewPassword)

	filter := bson.D{{Key: "username", Value: username}}
	update := bson.D{{Key: "$set", Value: bson.D{{Key: "password", Value: hashedPassword}}}}
	_, err = db.Collection("users").UpdateOne(context.Background(), filter, update)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Failed to reset password"})
	}

	return c.Status(200).JSON(fiber.Map{"message": "Password reset successfully"})
}
