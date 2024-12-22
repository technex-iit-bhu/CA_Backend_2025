package user

import (
	"CA_Backend/database"
	"CA_Backend/models"
	"CA_Backend/utils"
	"context"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func CreateUser(c *fiber.Ctx) error {
	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
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

	collection := db.Collection("users")

	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	user.CA_ID = utils.GenerateCAID(*user)
	user.Points = 0

	user.ReferralCode = utils.GetReferralCode(*user)

	if !utils.IsSafe(user.Password) {
		return c.Status(400).JSON(fiber.Map{
			"error":   "Password is not safe!!",
			"message": "Password must contain 8 characters, 1 uppercase & 1 lowercase letter, 1 number and 1 special character",
		})
	}

	var existingUser models.User
	filter := bson.D{{Key: "username", Value: user.Username}}
	err = collection.FindOne(ctx, filter).Decode(&existingUser)
	if err == nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Username already exists!!",
		})
	}

	filter = bson.D{{Key: "phoneNumber", Value: user.PhoneNumber}}
	err = collection.FindOne(ctx, filter).Decode(&existingUser)
	if err == nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Phone number already exists!!",
		})
	}

	user.Password = utils.HashPassword(user.Password)
	if r, err := collection.InsertOne(ctx, user); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error":   err.Error(),
			"message": "Failed to create user!!",
		})
	} else {
		return c.Status(201).JSON(fiber.Map{
			"id":      r.InsertedID,
			"message": "User created successfully",
		})
	}
}
