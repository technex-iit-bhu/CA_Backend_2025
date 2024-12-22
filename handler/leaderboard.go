package handler

import (
	"CA_Backend/database"
	"context"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetLeaderboard(c *fiber.Ctx) error {
	ctx := context.Background()
	db, err := database.Connect()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": err.Error()})
	}

	collection := db.Collection("users")

	findOptions := options.Find()
	findOptions.SetSort(bson.D{{Key: "points", Value: -1}})
	findOptions.SetLimit(20)

	cursor, err := collection.Find(ctx, bson.D{}, findOptions)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": err.Error()})
	}
	defer cursor.Close(ctx)

	var users []struct {
		Name   string `bson:"name"`
		CA_ID  string `bson:"ca_id"`
		Points int    `bson:"points"`
	}

	if err = cursor.All(ctx, &users); err != nil {
		return c.Status(500).JSON(fiber.Map{"message": err.Error()})
	}

	var leaderboard []fiber.Map
	for rank, user := range users {
		leaderboard = append(leaderboard, fiber.Map{
			"name":   user.Name,
			"ca_id":  user.CA_ID,
			"points": user.Points,
			"rank":   rank + 1,
		})
	}

	return c.Status(200).JSON(leaderboard)
}
