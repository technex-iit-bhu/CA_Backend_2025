package handler

import (
	"CA_Backend/database"
	"CA_Backend/models"
	"context"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserRank struct {
	Name   string `json:"name"`
	CA_ID  string `json:"ca_id"`
	Points int    `json:"points"`
	Rank   int    `json:"rank"`
}

func GetLeaderboard(c *fiber.Ctx) error {
	ctx := context.Background()
	db, err := database.Connect()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": err.Error()})
	}

	var users []models.User
	findOptions := options.Find().SetSort(
		bson.D{
			{Key: "points", Value: -1},
			{Key: "createdAt", Value: 1},
		},
	)

	cursor, err := db.Collection("users").Find(ctx, bson.D{}, findOptions)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": err.Error()})
	}
	if err = cursor.All(ctx, &users); err != nil {
		return c.Status(500).JSON(fiber.Map{"message": err.Error()})
	}
	defer cursor.Close(ctx)

	userRanks := make([]UserRank, len(users))
	for i, user := range users {
		userRanks[i] = UserRank{
			Name:   user.Name,
			CA_ID:  user.CA_ID,
			Points: user.Points,
		}
	}

	rank := 1
	for i := range userRanks {
		if i > 0 && userRanks[i].Points < userRanks[i-1].Points {
			rank = i + 1
		}
		userRanks[i].Rank = rank
	}

	return c.Status(200).JSON(userRanks)
}
