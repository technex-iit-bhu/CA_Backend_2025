package user

import (
	"CA_Backend/database"
	"CA_Backend/models"
	"CA_Backend/utils"
	"context"
	
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
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
	token := c.Get("Authorization")
	if len(token) > 7 { token = token[7:] } else { token = "" }
	username, err := utils.DeserialiseUser(token)
	if err != nil || updatedUser.Username != username{
		return c.Status(401).JSON(fiber.Map{"message": "Invalid user"})
	}

	user := new(models.User)
	if err := collection.FindOne(ctx, models.User{Username: updatedUser.Username}).Decode(user); err != nil {
		return c.Status(404).JSON(fiber.Map{"message": "User not found"})
	}

	if user.IsReferred && updatedUser.ReferralCode != "" {
		return c.Status(404).JSON(fiber.Map{"message": "Cannot update referral code again"})
	}

	if _, err := collection.UpdateOne(ctx, models.User{Username: updatedUser.Username}, bson.M{"$set": updatedUser}); err != nil {
		return c.Status(500).JSON(fiber.Map{"message": err.Error()})
	} else {
		return c.Status(200).JSON(fiber.Map{
			"message": "User updated Successfully",
			"user":    updatedUser,
		})
	}
}
