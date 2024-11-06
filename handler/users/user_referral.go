package user

import (
	"CA_Backend/database"
	"CA_Backend/models"
	"context"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func IncrementReferral(c *fiber.Ctx) error {
	ctx := context.Background()
	db, err := database.Connect()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": err.Error()})
	}

	type ReferralCode struct {
		ReferralCode string `json:"referral_code,omitempty" bson:"referral_code,omitempty" binding:"required"`
	}

	body := new(ReferralCode)
	if err := c.BodyParser(body); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error":   err.Error(),
			"message": "Failed to parse JSON Body",
		})
	}

	filter := bson.D{{Key: "referral_code", Value: body.ReferralCode}}

	update := bson.D{{Key: "$inc", Value: bson.D{{Key: "referral_count", Value: 1}}}}

	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)
	user := new(models.User)
	err = db.Collection("user").FindOneAndUpdate(ctx, filter, update, opts).Decode(user)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error":   err.Error(),
			"message": "Failed to update referral count",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Referral count incremented successfully",
		"user":    user,
	})
}
