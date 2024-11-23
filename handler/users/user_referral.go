package user

import (
	"CA_Backend/database"
	"CA_Backend/models"
	"context"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"strings"
	"time"
)

func SetReferral(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	db, err := database.Connect()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": err.Error()})
	}

	type Request struct {
		ReferralCode string `json:"referral_code,omitempty" bson:"referral_code,omitempty" binding:"required"`
		Username     string `json:"username,omitempty" bson:"username,omitempty" binding:"required"`
	}

	body := new(Request)
	if err := c.BodyParser(body); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error":   err.Error(),
			"message": "Failed to parse JSON Body",
		})
	}

	code := body.ReferralCode
	index := strings.Index(code, "_ca_")
	if index == -1 {
		return c.Status(400).JSON(fiber.Map{
			"error":   "Invalid Referral Code",
			"message": "Referral Code does not match the required format",
		})
	}
	ca_id := code[index+4:]
	filter := bson.D{{Key: "ca_id", Value: ca_id}}

	update := bson.D{{Key: "$inc", Value: bson.D{{Key: "referral_count", Value: 1}}}}
	user := new(models.User)
	err = db.Collection("user").FindOneAndUpdate(
		ctx,
		filter,
		update,
		options.FindOneAndUpdate().SetReturnDocument(options.After),
	).Decode(user)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error":   err.Error(),
			"message": "Failed to update referral count",
		})
	}

	currentUserFilter := bson.D{{Key: "username", Value: body.Username}}
	updateReferred := bson.D{{Key: "$set", Value: bson.D{{Key: "is_referred", Value: true}}}}
	_, err = db.Collection("user").UpdateOne(ctx, currentUserFilter, updateReferred)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error":   err.Error(),
			"message": "Failed to set Referral status",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Updated Referral Successfully",
		"user":    user,
	})
}
